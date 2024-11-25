package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	deployHandler "general_spider_controll_panel/handler/deploy"
	"general_spider_controll_panel/types/models"
	"general_spider_controll_panel/utils"
	configView "general_spider_controll_panel/view/config"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")
var spider = "general_engine"

var (
	TestRun = make(map[string][]byte)
	mu      sync.Mutex
)

func GET(w http.ResponseWriter, r *http.Request) {

	configView.Main("config maker").Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	configID := uuid.New()
	baseUrl := r.Form.Get("base-url")
	name := r.Form.Get("name")
	configType := r.Form.Get("configType")
	description := r.Form.Get("description")
	jsonData := r.Form.Get("jsonData")

	parse, err := url.Parse(baseUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "test-config":
			jobID := utils.GetMD5Hash(jsonData)
			go deployHandler.RunSpider(scrapydURL, parse.Host, spider, map[string]string{
				"preview":        "yes",
				"preview_config": base64.StdEncoding.EncodeToString([]byte(jsonData)),
				"jobid":          jobID,
			})
			fmt.Println("Hash:", jobID)

			data, err := waitForResult(jobID, 100, 1*time.Second)
			if err != nil {
				http.Error(w, "<pre>Cannot get the data in time, please try again</pre>", http.StatusRequestTimeout)
				return
			}
			formattedData, err := formatJSON(data)
			fmt.Fprintf(w, "<pre>%s</pre>", formattedData)
			return
		case "save-config":
			if name == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = app.Server.Database.CreateConfig(&models.Config{
				ID:          configID,
				Domain:      parse.Host,
				Name:        name,
				Type:        configType,
				Description: description,
				Data:        []byte(jsonData),
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Hx-Redirect", "/deploy")
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

func waitForResult(jobID string, maxRetry int, delay time.Duration) ([]byte, error) {
	for retry := 0; retry < maxRetry; retry++ {
		mu.Lock()
		data, ok := TestRun[jobID]
		mu.Unlock()

		if ok {
			return data, nil
		}
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("timed out waiting for result")
}

func formatJSON(data []byte) (string, error) {
	var prettyJSON map[string]interface{}
	err := json.Unmarshal(data, &prettyJSON)
	if err != nil {
		return "", err
	}

	indentedJSON, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		return "", err
	}

	return string(indentedJSON), nil
}
