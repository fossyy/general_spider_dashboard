package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type ScrapydStruct struct {
	ScrapydURL string
	Version    string
	Spider     string
}

type ScrapydProjectsResponse struct {
	Status   string   `json:"status"`
	Projects []string `json:"projects"`
}

func (sc *ScrapydStruct) RunSpider(project string, additionalArgs map[string]string) (string, error) {
	data := url.Values{}
	data.Set("project", project)
	data.Add("spider", sc.Spider)

	for key, value := range additionalArgs {
		data.Add(key, value)
	}

	url := fmt.Sprintf("%s/schedule.json", sc.ScrapydURL)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("run failed with status: %s", resp.Status)
	}

	var scrapydResp types.ScrapydResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	if scrapydResp.Status == "error" {
		fmt.Println("error : ", scrapydResp.Message)
		pattern := `^project '.*' not found$`
		re := regexp.MustCompile(pattern)
		if re.MatchString(scrapydResp.Message) {
			err := sc.UploadEgg(project)
			if err != nil {
				return "", err
			}
			return sc.RunSpider(project, additionalArgs)
		}
		return "", fmt.Errorf("%s", scrapydResp.Message)
	}
	return scrapydResp.Jobid, nil
}

func (sc *ScrapydStruct) GetAllProjects() ([]string, error) {
	url := fmt.Sprintf("%s/listprojects.json", sc.ScrapydURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch projects from Scrapyd: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch projects, status: %s", resp.Status)
	}

	var projectsResponse ScrapydProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&projectsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if projectsResponse.Status != "ok" {
		return nil, fmt.Errorf("unexpected response status: %s", projectsResponse.Status)
	}

	return projectsResponse.Projects, nil
}

func (sc *ScrapydStruct) UploadEgg(project string) error {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	if err := writer.WriteField("project", project); err != nil {
		return fmt.Errorf("failed to add project field: %w", err)
	}

	if err := writer.WriteField("version", sc.Version); err != nil {
		return fmt.Errorf("failed to add version field: %w", err)
	}

	part, err := writer.CreateFormFile("egg", "general.egg")
	if err != nil {
		return fmt.Errorf("failed to add egg file field: %w", err)
	}
	if _, err := io.Copy(part, bytes.NewReader(egg)); err != nil {
		return fmt.Errorf("failed to copy egg file data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/addversion.json", sc.ScrapydURL)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status: %s", resp.Status)
	}

	Server.Logger.Println("Egg uploaded successfully!")
	return nil
}

func (sc *ScrapydStruct) UploadEggToAllProjects(database types.Database) error {
	dbProjects, err := database.GetDomains()
	if err != nil {
		return fmt.Errorf("failed to get projects: %w", err)
	}
	scrapydProjects, err := sc.GetAllProjects()
	if err != nil {
		return fmt.Errorf("failed to get projects: %w", err)
	}
	var projects []string
	exists := make(map[string]bool)

	for _, project := range dbProjects {
		if !exists[project] {
			exists[project] = true
			projects = append(projects, project)
		}
	}

	for _, project := range scrapydProjects {
		if !exists[project] {
			exists[project] = true
			projects = append(projects, project)
		}
	}

	for _, project := range projects {
		Server.Logger.Printf("Uploading egg to project: %s\n", project)
		if err := sc.UploadEgg(project); err != nil {
			Server.Logger.Printf("Failed to upload egg to project %s: %v\n", project, err)
		}
	}

	return nil
}

func (sc *ScrapydStruct) GetRunningSpiders(project string) (*types.ScrapydResponseGetingSpiders, error) {
	url := fmt.Sprintf("%s/listjobs.json?project=%s", sc.ScrapydURL, project)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var scrapydResp types.ScrapydResponseGetingSpiders
	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return &scrapydResp, nil
}

func (sc *ScrapydStruct) GetSpider(jobID string, projects []string) (*types.Spider, error) {
	var data *types.Spider

	for _, project := range projects {
		url := fmt.Sprintf("%s/listjobs.json?project=%s", sc.ScrapydURL, project)

		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to send HTTP request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		var scrapydResp types.ScrapydResponseGetingSpiders
		err = json.Unmarshal(body, &scrapydResp)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
		}
		for _, runningSpider := range scrapydResp.Running {
			if runningSpider.Id == jobID {
				data = &types.Spider{
					Id:        runningSpider.Id,
					Spider:    runningSpider.Spider,
					Pid:       runningSpider.Pid,
					ItemsUrl:  runningSpider.ItemsUrl,
					LogUrl:    runningSpider.LogUrl,
					Project:   runningSpider.Project,
					StartTime: runningSpider.StartTime,
					EndTime:   "Still Running",
					Status:    "Running",
					NodeName:  scrapydResp.NodeName,
				}
			}
		}

		if data == nil {
			for _, pendingSpider := range scrapydResp.Pending {
				if pendingSpider.Id == jobID {
					data = &types.Spider{
						Id:        pendingSpider.Id,
						Spider:    pendingSpider.Spider,
						Pid:       0,
						ItemsUrl:  "",
						LogUrl:    "",
						Project:   pendingSpider.Project,
						StartTime: "Pending",
						EndTime:   "Not Running Yet",
						Status:    "Pending",
						NodeName:  scrapydResp.NodeName,
					}
				}
			}
		}

		if data == nil {
			for _, finishedSpider := range scrapydResp.Finished {
				if finishedSpider.Id == jobID {
					data = &types.Spider{
						Id:        finishedSpider.Id,
						Spider:    finishedSpider.Spider,
						Pid:       0,
						ItemsUrl:  finishedSpider.ItemsUrl,
						LogUrl:    finishedSpider.LogUrl,
						Project:   finishedSpider.Project,
						StartTime: finishedSpider.StartTime,
						EndTime:   finishedSpider.EndTime,
						Status:    "Finished",
						NodeName:  scrapydResp.NodeName,
					}
				}
			}
		}
		if data != nil {
			break
		}
	}

	return data, nil
}

func (sc *ScrapydStruct) FetchScrapydLog(project, jobID string, maxlen int) ([]string, error) {
	url := fmt.Sprintf("%s/spiderlogs.json?project=%s&jobid=%s&maxlen=%d", sc.ScrapydURL, project, jobID, maxlen)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var scrapydResp []string
	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return []string{"Error when getting logs : " + err.Error()}, nil
	}

	return scrapydResp, nil
}

func (sc *ScrapydStruct) GetSpiderUsage(project, jobID string) (*types.SpiderUsage, error) {
	url := fmt.Sprintf("%s/spiderstatus.json?project=%s&jobid=%s", sc.ScrapydURL, project, jobID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var scrapydResp types.SpiderUsage
	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return &scrapydResp, nil
}

func (sc *ScrapydStruct) GetLogsAndResults() (*types.ScrapydLogsAndResults, error) {
	url := fmt.Sprintf("%s/spiderstorage.json", sc.ScrapydURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var scrapydResp types.ScrapydLogsAndResults
	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return &scrapydResp, nil
}

func (sc *ScrapydStruct) GetLog(project, job_id string) ([]byte, error) {
	url := fmt.Sprintf("%s/spiderdownloadlog.json?project=%s&job_id=%s", sc.ScrapydURL, project, job_id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func (sc *ScrapydStruct) GetResult(project, job_id string) ([]byte, error) {
	url := fmt.Sprintf("%s/spiderdownloadresult.json?project=%s&job_id=%s", sc.ScrapydURL, project, job_id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get running spiders, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
