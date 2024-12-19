package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types/models"
	"general_spider_controll_panel/utils"
	configView "general_spider_controll_panel/view/config"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	TestRun = make(map[string][]byte)
	mu      sync.Mutex
)

func GET(w http.ResponseWriter, r *http.Request) {
	domains, err := app.Server.Database.GetDomainsWithSchema()
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	configView.Main("config maker", domains).Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
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
			if parse.Host == "" || parse.Scheme == "" {
				fmt.Fprintf(w, fmt.Sprintf("<pre>Invalid URL with format of %s://%s:%s, Please check the format and try again.</pre>", strings.TrimSpace(parse.Scheme), strings.TrimSpace(parse.Host), strings.TrimSpace(parse.Port())))
				w.WriteHeader(http.StatusOK)
				return
			}
			jobID := fmt.Sprintf("preview_%s", utils.GetMD5Hash(jsonData))
			proxies, err := app.Server.Database.GetActiveProxies()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if len(proxies) == 0 {
				fmt.Fprintln(w, "<pre>No available proxies. Please add a proxy and try again.</pre>")
				return
			}
			preview_proxies, err := json.Marshal([]string{fmt.Sprintf("%s://%s:%s", strings.TrimSpace(proxies[0].Protocol), strings.TrimSpace(proxies[0].Address), strings.TrimSpace(proxies[0].Port))})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			go func() {
				_, err := app.Server.Scrapyd.RunSpider(parse.Host, map[string]string{
					"preview":         "yes",
					"preview_config":  base64.StdEncoding.EncodeToString([]byte(jsonData)),
					"preview_proxies": base64.StdEncoding.EncodeToString(preview_proxies),
					"jobid":           jobID,
				})
				if err != nil {
					mu.Lock()
					errorJson, _ := json.Marshal(map[string]string{
						"error": err.Error(),
					})
					TestRun[jobID] = errorJson
					mu.Unlock()
				}
			}()

			data, err := waitForResult(r.Context(), jobID, 30, 1*time.Second)
			if err != nil {
				fmt.Fprintln(w, "<pre>Cannot get the data in time, please try again</pre>")
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
				ID:             configID,
				Domain:         parse.Host,
				DomainProtocol: parse.Scheme,
				Name:           name,
				Type:           configType,
				Description:    description,
				Data:           []byte(jsonData),
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

func waitForResult(ctx context.Context, jobID string, maxRetry int, delay time.Duration) ([]byte, error) {
	for retry := 0; retry < maxRetry; retry++ {
		select {
		case <-ctx.Done():
			app.Server.Logger.Println(fmt.Errorf("operation canceled: %w", ctx.Err()))
			return nil, fmt.Errorf("operation canceled: %w", ctx.Err())
		default:
		}
		mu.Lock()
		data, ok := TestRun[jobID]
		mu.Unlock()

		if ok {
			return data, nil
		}
		app.Server.Logger.Println("Waiting for job", jobID, "to finish")
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
