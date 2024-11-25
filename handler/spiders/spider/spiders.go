package HandlerSpiders

import (
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	spiderView "general_spider_controll_panel/view/spider/spider"
	"io"
	"net/http"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")
var version = "1.0"

func GET(w http.ResponseWriter, r *http.Request) {
	project := r.PathValue("project")
	err := ensureProjectAndUpload(scrapydURL, project, version)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println(err.Error())
		return
	}
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-spiders":
			spiders, err := GetRunningSpiders(scrapydURL, project)
			if err != nil {
				return
			}
			spiderView.GetSpider(spiders).Render(r.Context(), w)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	spiderView.Main("Spider Page").Render(r.Context(), w)
	return
}

func GetRunningSpiders(scrapydURL, project string) (*types.ScrapydResponseGetingSpiders, error) {
	url := fmt.Sprintf("%s/listjobs.json?project=%s", scrapydURL, project)

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

func ensureProjectAndUpload(scrapydURL, project, version string) error {
	projects, err := app.GetAllProjects(scrapydURL)
	if err != nil {
		return fmt.Errorf("failed to fetch projects: %w", err)
	}

	projectExists := false
	for _, existingProject := range projects {
		if existingProject == project {
			projectExists = true
			break
		}
	}

	if !projectExists {
		fmt.Printf("Project '%s' does not exist. Creating it by uploading the egg.\n", project)
	}

	return app.UploadEgg(scrapydURL, project, version)
}
