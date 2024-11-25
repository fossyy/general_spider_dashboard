package handlerSpiderDetails

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	spiderDetailsView "general_spider_controll_panel/view/spider/details"
	"github.com/shirou/gopsutil/v3/process"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")
var spider = "general_engine"

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
	id := r.PathValue("id")
	project := r.PathValue("project")
	detail, err := getSpiderDetail(id, project)
	if err != nil && detail != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if detail == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "http-status":
			codes, err := countStatusCodesFromLog(project, id, statusCode)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			spiderDetailsView.HttpStatusUI(codes).Render(r.Context(), w)
			return
		case "performance-metrics":
			spiderDetailsView.PerformanceMatricsUI(detail).Render(r.Context(), w)
			return
		case "spider-status":
			spiderDetailsView.SpiderStatusUI(detail).Render(r.Context(), w)
			return
		case "spider-name":
			spiderDetailsView.SpiderNameUI(fmt.Sprintf("%s_%s", detail.Name, detail.Id)).Render(r.Context(), w)
			return
		case "spider-logs":
			spiderDetailsView.SpiderLogsUI(detail).Render(r.Context(), w)
			return
		case "spider-actions":
			if detail.Status == "Running" {
				spiderDetailsView.SpiderActionsUI(detail.Name, detail.Id).Render(r.Context(), w)
				return
			}
			w.Write([]byte(""))
			return
		default:
			http.NotFound(w, r)
			return
		}
	}
	spiderDetailsView.Main("Spider page").Render(r.Context(), w)
	return
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project := r.PathValue("project")
	detail, err := getSpiderDetail(id, project)
	if err != nil && detail != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if detail == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = stopSpider(project, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Refresh", "true")
	return
}

func getSpiderDetail(jobID, project string) (*types.SpiderDetail, error) {
	spiders, err := getSpider(jobID, scrapydURL, project)
	if err != nil {
		fmt.Println("error 1 ", err.Error())
		return nil, err
	}

	if spiders == nil {
		fmt.Println("error 2 ", err.Error())
		return nil, fmt.Errorf("spider not found with JobID: %s", jobID)
	}

	pages, err := countCrawledPagesFromLog(project, jobID)
	if err != nil {
		fmt.Println("error 3 ", err.Error())
		return nil, err
	}

	tailLog, err := readLastLinesFromLog(project, jobID, 40)
	if err != nil {
		fmt.Println("error 4 ", err.Error())
		tailLog = []string{"Null"}
	}

	var spiderDetail types.SpiderDetail
	if spiders.Status == "Running" {
		var cpuPercent float64
		var memInfoRSS uint64
		if spiders.Pid < 100 {
			cpuPercent = float64(0)
			memInfoRSS = uint64(0)
		} else {
			proc, _ := process.NewProcess(int32(spiders.Pid))
			cpuPercent, _ = proc.CPUPercent()
			memInfo, _ := proc.MemoryInfo()
			memInfoRSS = memInfo.RSS
		}

		spiderDetail.Cpu = fmt.Sprintf("%.2f%%", cpuPercent)
		spiderDetail.Mem = memInfoRSS
		spiderDetail.NodeName = spiders.NodeName
		spiderDetail.CrawledCount = pages
		spiderDetail.PID = spiders.Pid
		spiderDetail.Log = tailLog
		spiderDetail.Name = spiders.Spider
		spiderDetail.Status = spiders.Status
		spiderDetail.Id = spiders.Id
		spiderDetail.StartTime = spiders.StartTime
		spiderDetail.EndTime = spiders.EndTime
	} else {
		spiderDetail.Cpu = "0"
		spiderDetail.Mem = 0
		spiderDetail.NodeName = spiders.NodeName
		spiderDetail.CrawledCount = pages
		spiderDetail.PID = 0
		spiderDetail.Log = tailLog
		spiderDetail.Name = spiders.Spider
		spiderDetail.Status = spiders.Status
		spiderDetail.Id = spiders.Id
		spiderDetail.StartTime = spiders.StartTime
		spiderDetail.EndTime = spiders.EndTime
	}
	return &spiderDetail, nil
}

func getSpider(jobID, scrapydURL, project string) (*types.Spider, error) {
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

	fmt.Println("body : ", string(body))
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
	return data, nil
}

func fetchScrapydLog(project, jobID string) (string, error) {
	logFilePath := fmt.Sprintf("logs/%s/%s/log.log", project, jobID)
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		return "Logs path is not valid \n", nil
	}
	file, err := os.Open(logFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read log file: %w", err)
	}

	return string(content), nil
}

func readLastLinesFromLog(project, jobID string, numLines int) ([]string, error) {
	logContent, err := fetchScrapydLog(project, jobID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch log: %w", err)
	}

	lines := make([]string, 0, numLines)
	var currentLine string
	lineCount := 0
	reader := bufio.NewReader(strings.NewReader(logContent))
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
	return lines, nil
}

func countCrawledPagesFromLog(project, jobID string) (uint64, error) {
	logContent, err := fetchScrapydLog(project, jobID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch log: %w", err)
	}
	re := regexp.MustCompile(`\|\s*(\d+)`)

	var crawlCount uint64 = 0

	reader := bufio.NewReader(strings.NewReader(logContent))
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return crawlCount, err
		}
		if len(line) == 0 {
			break
		}

		line = strings.TrimSpace(line)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			crawlCountInt, err := strconv.Atoi(matches[1])
			if err != nil {
				return crawlCount, fmt.Errorf("failed to parse crawled pages count: %w", err)
			}
			crawlCount = uint64(crawlCountInt)
		}

		if err != nil {
			break
		}
	}

	return crawlCount, nil
}

func countStatusCodesFromLog(project, jobID string, statusCodeMap map[int]string) ([]*types.StatusCode, error) {
	logContent, err := fetchScrapydLog(project, jobID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch log: %w", err)
	}
	re := regexp.MustCompile(`\[\w+]\s*\|\s*(\{.*})$`)
	var statusCounts []types.StatusCode

	reader := bufio.NewReader(strings.NewReader(logContent))
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return nil, err
		}
		if len(line) == 0 {
			break
		}

		line = strings.TrimSpace(line)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			statusCodeDataStr := matches[1]
			statusCodeDataStr = strings.Trim(statusCodeDataStr, "{}")
			pairs := strings.Split(statusCodeDataStr, ",")
			for _, pair := range pairs {
				kv := strings.Split(pair, ":")
				if len(kv) == 2 {
					codeStr := strings.TrimSpace(kv[0])
					countStr := strings.TrimSpace(kv[1])
					code, err := strconv.Atoi(codeStr)
					if err != nil {
						continue
					}
					count, err := strconv.Atoi(countStr)
					if err != nil {
						continue
					}

					if message, exists := statusCodeMap[code]; exists {
						found := false
						for i := range statusCounts {
							if statusCounts[i].Code == uint(code) {
								statusCounts[i].Count = uint(count)
								found = true
								break
							}
						}
						if !found {
							statusCounts = append(statusCounts, types.StatusCode{
								Code:      uint(code),
								Detail:    message,
								Count:     uint(count),
								BaseGroup: getBaseGroup(code),
							})
						}
					}
				}
			}
		}

		if err != nil {
			break
		}
	}

	var filteredStatusCounts []*types.StatusCode
	for _, status := range statusCounts {
		if status.Count > 0 {
			filteredStatusCounts = append(filteredStatusCounts, &status)
		}
	}

	return filteredStatusCounts, nil
}

func stopSpider(project, jobID string) error {
	addr := fmt.Sprintf("%s/cancel.json", scrapydURL)
	data := url.Values{}
	data.Set("project", project)
	data.Set("job", jobID)

	req, err := http.NewRequest("POST", addr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	return nil
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
