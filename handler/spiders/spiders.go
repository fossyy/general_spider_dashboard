package handlerSpiders

import (
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	spiderView "general_spider_controll_panel/view/spider"
	"io"
	"net/http"
)

var scrapydURL = "http://localhost:6800"
var project = "general"
var spider = "general_engine"
var version = "1.0"
var eggPath = "general.egg"

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-spiders":
			spiders, err := getRunningSpiders(scrapydURL, project)
			if err != nil {
				return
			}
			spiderView.GetSpider(spiders).Render(r.Context(), w)
			return
		default:
			spiderView.Main("Spider Page").Render(r.Context(), w)
			return
		}
	}
	spiderView.Main("Spider Page").Render(r.Context(), w)
	return
}

func getRunningSpiders(scrapydURL, project string) (*types.ScrapydResponseGetingSpiders, error) {
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
