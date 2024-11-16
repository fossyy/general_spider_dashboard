package handlerSpiderDetails

import (
	"bufio"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	spiderDetailsView "general_spider_controll_panel/view/spider/details"
	"github.com/shirou/gopsutil/v3/process"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var scrapydURL = "http://localhost:6800"
var project = "general"
var spider = "general_engine"
var version = "1.0"
var eggPath = "general.egg"

var statusCode = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a Teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		id := r.PathValue("id")
		details, err := getSpidersDetails(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		switch r.URL.Query().Get("action") {
		case "http-status":
			spiderDetailsView.HttpStatusUI(details).Render(r.Context(), w)
			return
		case "performance-metrics":
			spiderDetailsView.PerformanceMatricsUI(details).Render(r.Context(), w)
			return
		case "spider-status":
			spiderDetailsView.SpiderStatusUI(details).Render(r.Context(), w)
			return
		case "spider-name":
			spiderDetailsView.SpiderNameUI(fmt.Sprintf("%s_%s", details.Spider.Spider, details.Spider.Id)).Render(r.Context(), w)
			return
		case "spider-logs":
			spiderDetailsView.SpiderLogsUI(details).Render(r.Context(), w)
			return
		default:
			http.NotFound(w, r)
			return
		}
	}
	spiderDetailsView.Main("Spider page").Render(r.Context(), w)
}

func getSpidersDetails(jobID string) (*types.SpiderDetails, error) {
	url := fmt.Sprintf("%s/status.json?job=%s", scrapydURL, jobID)

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

	var scrapydResp types.SpiderDetails
	err = json.Unmarshal(body, &scrapydResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}

	spiders, err := getRunningSpider(jobID, scrapydURL, project)
	if err != nil {
		return nil, err
	}
	pages, err := countCrawledPages("/home/bagas/Documents/scrapy_engine" + spiders.LogUrl)
	if err != nil {
		return nil, err
	}
	if scrapydResp.Currstate == "running" {
		proc, _ := process.NewProcess(int32(spiders.Pid))
		cpuPercent, _ := proc.CPUPercent()
		memInfo, _ := proc.MemoryInfo()
		scrapydResp.Detail.Cpu = fmt.Sprintf("%.2f%%", cpuPercent)
		scrapydResp.Detail.Mem = memInfo.RSS
		scrapydResp.Detail.NodeName = scrapydResp.NodeName
		scrapydResp.Detail.CrawledCount = pages
	} else {
		scrapydResp.Detail.Cpu = "0"
		scrapydResp.Detail.Mem = 0
		scrapydResp.Detail.NodeName = scrapydResp.NodeName
		scrapydResp.Detail.CrawledCount = pages
	}

	scrapydResp.Spider = spiders
	tailLog, err := readLastLines("/home/bagas/Documents/scrapy_engine"+spiders.LogUrl, 40)
	if err != nil {
		tailLog = []string{"Null"}
	}
	scrapydResp.Log = tailLog
	codes, err := countStatusCodes("/home/bagas/Documents/scrapy_engine"+spiders.LogUrl, statusCode)
	scrapydResp.Detail.CrawledDetail = codes
	return &scrapydResp, nil
}

func getRunningSpider(jobID, scrapydURL, project string) (*types.Spider, error) {
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

	var data *types.Spider
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
				}
			}
		}
	}

	return data, nil
}

func readLastLines(filePath string, numLines int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	lines := make([]string, 0, numLines)

	reader := bufio.NewReader(file)

	var currentLine string

	lineCount := 0
	for {
		currentLine, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		if lineCount >= numLines {
			lines = lines[1:]
		}
		lines = append(lines, currentLine)
		lineCount++
	}

	if err != nil && err.Error() != "EOF" {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

func countCrawledPages(filePath string) (uint64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	re := regexp.MustCompile(`Crawled (\d+) pages`)

	totalPages := uint64(0)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return 0, fmt.Errorf("failed to read file: %w", err)
		}
		if err == nil && len(line) == 0 {
			break
		}

		line = strings.TrimSpace(line)

		matches := re.FindStringSubmatch(line)
		if len(matches) == 2 {
			crawledPages, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, fmt.Errorf("failed to parse crawled pages count: %w", err)
			}
			totalPages += uint64(crawledPages)
		}

		if err != nil {
			break
		}
	}

	return totalPages, nil
}

func countStatusCodes(filePath string, statusCodeMap map[int]string) ([]types.StatusCode, error) {
	var statusCounts []types.StatusCode
	codeIndexMap := make(map[int]int)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	re := regexp.MustCompile(`Crawled \((\d{3})\)`)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return nil, err
		}
		if err == nil && len(line) == 0 {
			break
		}

		line = strings.TrimSpace(line)

		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			codeStr := matches[1]
			code, _ := strconv.Atoi(codeStr)

			if message, exists := statusCodeMap[code]; exists {
				if index, found := codeIndexMap[code]; found {
					statusCounts[index].Count++
				} else {
					statusCounts = append(statusCounts, types.StatusCode{
						Code:      uint(code),
						Detail:    message,
						Count:     1,
						BaseGroup: getBaseGroup(code),
					})
					codeIndexMap[code] = len(statusCounts) - 1
				}
			}
		}
		if err != nil {
			break
		}
	}

	if err != nil && err.Error() != "EOF" {
		return nil, err
	}

	var filteredStatusCounts []types.StatusCode
	for _, status := range statusCounts {
		if status.Count > 0 {
			filteredStatusCounts = append(filteredStatusCounts, status)
		}
	}

	return filteredStatusCounts, nil
}

func getBaseGroup(code int) string {
	switch {
	case code >= 100 && code < 200:
		return "1xx"
	case code >= 200 && code < 300:
		return "2xx"
	case code >= 300 && code < 400:
		return "3xx"
	case code >= 400 && code < 500:
		return "4xx"
	case code >= 500 && code < 600:
		return "5xx"
	default:
		return "Unknown"
	}
}
