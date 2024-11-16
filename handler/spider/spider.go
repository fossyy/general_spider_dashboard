package handlerSpider

import (
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var scrapydURL = "http://localhost:6800"
var project = "general"
var spider = "general_engine"
var version = "1.0"
var eggPath = "general.egg"

func POST(w http.ResponseWriter, r *http.Request) {
	baseDir := "configs"
	var configsList []types.Config

	dirs, err := os.ReadDir(baseDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			baseURL := "http://" + dir.Name()

			dirPath := filepath.Join(baseDir, dir.Name())
			files, err := os.ReadDir(dirPath)
			if err != nil {
				fmt.Println("Error reading directory:", err)
				return
			}

			var jsonFiles []string
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
					fileName := strings.TrimSuffix(file.Name(), ".json")
					jsonFiles = append(jsonFiles, fileName)
				}
			}

			configsList = append(configsList, types.Config{
				BaseURL: baseURL,
				Configs: jsonFiles,
			})
		}
	}

	json.NewEncoder(w).Encode(configsList)
}

func handleRunSpider(w http.ResponseWriter, r *http.Request) {
	config := r.PathValue("config")
	additionalArgs := map[string]string{
		"config_path": config,
	}
	jobID, err := runSpider(scrapydURL, project, spider, additionalArgs)
	if err != nil {
		fmt.Fprintf(w, "Error running spider: %v\n", err)
	} else {
		fmt.Fprintf(w, "Spider scheduled successfully! Job ID: %s\n", jobID)
	}
	fmt.Fprintf(w, "Running with id : %s", jobID)
}

func runSpider(scrapydURL, project, spider string, additionalArgs map[string]string) (string, error) {
	data := url.Values{}
	data.Set("project", project)
	data.Set("spider", spider)

	for key, value := range additionalArgs {
		data.Set(key, value)
	}

	url := fmt.Sprintf("%s/schedule.json", scrapydURL)
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

	fmt.Println(string(body))
	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	return scrapydResp.Jobid, nil
}
