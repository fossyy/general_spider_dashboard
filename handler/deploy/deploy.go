package deployHandler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	deployView "general_spider_controll_panel/view/deploy"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")
var spider = "general_engine"

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-proxies":
			proxies, err := app.Server.Database.GetProxies()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			deployView.ProxiesUI(proxies).Render(r.Context(), w)
			return
		case "get-configs":
			domain := r.URL.Query().Get("domainSelect")
			configs, err := app.Server.Database.GetConfigNameAndIDByDomain(domain)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			deployView.ConfigListUI(configs).Render(r.Context(), w)
			return
		case "get-additional-output-settings":
			outputDest := r.URL.Query().Get("outputDestination")
			if outputDest == "kafka" {
				deployView.KafkaSettingsUI().Render(r.Context(), w)
				return
			} else {
				w.Write([]byte(""))
				return
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	domains, err := app.Server.Database.GetDomains()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	deployView.Main("Deploy page", domains).Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	configID := r.Form.Get("configSelect")
	outputDest := r.Form.Get("outputDestination")
	kafkaBrokers := r.Form.Get("kafkaBrokers")
	kafkaTopic := r.Form.Get("kafkaTopic")
	proxies := r.Form.Get("selectedProxies")
	proxiesJSON, _ := json.Marshal(strings.Split(proxies, ","))
	additionalArgs := map[string]string{
		"config_id":  configID,
		"output_dst": outputDest,
		"proxies":    base64.StdEncoding.EncodeToString(proxiesJSON),
	}
	if kafkaBrokers != "" {
		additionalArgs["kafka_server"] = kafkaBrokers
	}
	if kafkaTopic != "" {
		additionalArgs["kafka_topic"] = kafkaTopic
	}
	project := r.Form.Get("domainSelect")
	_, err = RunSpider(scrapydURL, project, spider, additionalArgs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("proxies : ", base64.StdEncoding.EncodeToString(proxiesJSON))
	w.Header().Set("Hx-Redirect", fmt.Sprintf("/spiders/%s", project))
	w.WriteHeader(http.StatusCreated)
}

func RunSpider(scrapydURL, project, spider string, additionalArgs map[string]string) (string, error) {
	data := url.Values{}
	data.Set("project", project)
	data.Add("spider", spider)

	for key, value := range additionalArgs {
		data.Add(key, value)
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
