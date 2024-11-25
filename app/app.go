package app

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	"io"
	"mime/multipart"
	"net/http"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")
var version = "1.0"

var Server *App

//go:embed general.egg
var egg []byte

type App struct {
	http.Server
	Database types.Database
}

type ScrapydProjectsResponse struct {
	Status   string   `json:"status"`
	Projects []string `json:"projects"`
}

func GetAllProjects(scrapydURL string) ([]string, error) {
	url := fmt.Sprintf("%s/listprojects.json", scrapydURL)
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

func UploadEgg(scrapydURL, project, version string) error {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	if err := writer.WriteField("project", project); err != nil {
		return fmt.Errorf("failed to add project field: %w", err)
	}

	if err := writer.WriteField("version", version); err != nil {
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

	url := fmt.Sprintf("%s/addversion.json", scrapydURL)
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

	fmt.Println("Egg uploaded successfully!")
	return nil
}

func UploadEggToAllProjects(scrapydURL, version string, database types.Database) error {
	projects, err := database.GetDomains()
	if err != nil {
		return fmt.Errorf("failed to get projects: %w", err)
	}

	for _, project := range projects {
		fmt.Printf("Uploading egg to project: %s\n", project)
		if err := UploadEgg(scrapydURL, project, version); err != nil {
			fmt.Printf("Failed to upload egg to project %s: %v\n", project, err)
		}
	}

	return nil
}

func NewApp(addr string, handler http.Handler, database types.Database) *App {
	if err := UploadEggToAllProjects(scrapydURL, version, database); err != nil {
		panic("Error: " + err.Error())
	}
	return &App{
		Server: http.Server{
			Addr:    addr,
			Handler: handler,
		},
		Database: database,
	}
}
