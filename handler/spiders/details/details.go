package handlerSpiderDetails

import (
	"bytes"
	"errors"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	spiderDetailsView "general_spider_controll_panel/view/spider/details"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	detail, err := GetSpiderDetail(id, project)
	if err != nil && detail != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if detail == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log, err := countStatusCodesFromLog(detail.Log, statusCode)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	spiderDetailsView.Main("Spider page", detail, log).Render(r.Context(), w)
	return
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project := r.PathValue("project")
	detail, err := GetSpiderDetail(id, project)
	if err != nil && detail != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if detail == nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = stopSpider(project, id)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	max_retry := 30
	retry := 0
	for {
		if retry >= max_retry {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stopSpiderStatus, _ := app.Server.Scrapyd.GetSpider(id, []string{project})
		if stopSpiderStatus.Status != "Finished" {
			time.Sleep(1 * time.Second)
			retry += 1
			continue
		} else {
			proxies, err := app.Server.Database.GetProxiesByJobID(id)
			if err != nil {
				if errors.Is(err, gorm.ErrEmptySlice) {
					w.Header().Set("HX-Redirect", fmt.Sprintf("/spiders/%s", project))
					return
				}
				app.Server.Logger.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			for _, proxy := range proxies {
				app.Server.Tools.CheckProxy(proxy)
			}
			break
		}
	}
	w.Header().Set("HX-Refresh", "true")
	return
}

func GetSpiderDetail(jobID, project string) (*types.SpiderDetail, error) {
	spiders, err := app.Server.Scrapyd.GetSpider(jobID, []string{project})
	if err != nil {
		app.Server.Logger.Println(err)
		return nil, err
	}

	if spiders == nil {
		return nil, fmt.Errorf("spider not found with JobID: %s", jobID)
	}

	tailLog, err := app.Server.Scrapyd.FetchScrapydLog(project, jobID, 100)
	if err != nil {
		app.Server.Logger.Println(err)
		tailLog = []string{"Null"}
	}

	pages, err := countCrawledPagesFromLog(tailLog)
	if err != nil {
		app.Server.Logger.Println(err)
		return nil, err
	}

	var spiderDetail types.SpiderDetail
	if spiders.Status == "Running" {
		usage, err := app.Server.Scrapyd.GetSpiderUsage(project, jobID)
		if err != nil {
			app.Server.Logger.Println(err)
			return nil, err
		}
		spiderDetail.Cpu = fmt.Sprintf("%.2f%%", usage.Usage.Cpu)
		spiderDetail.Mem = fmt.Sprintf("%.2f MB", usage.Usage.Memory)
		spiderDetail.NodeName = spiders.NodeName
		spiderDetail.CrawledCount = pages
		spiderDetail.PID = spiders.Pid
		spiderDetail.Log = tailLog
		spiderDetail.Name = spiders.Spider
		spiderDetail.Status = spiders.Status
		spiderDetail.Id = spiders.Id
		spiderDetail.StartTime = spiders.StartTime
		spiderDetail.EndTime = spiders.EndTime
		spiderDetail.Project = project
	} else {
		spiderDetail.Cpu = "0"
		spiderDetail.Mem = "0"
		spiderDetail.NodeName = spiders.NodeName
		spiderDetail.CrawledCount = pages
		spiderDetail.PID = 0
		spiderDetail.Log = tailLog
		spiderDetail.Name = spiders.Spider
		spiderDetail.Status = spiders.Status
		spiderDetail.Id = spiders.Id
		spiderDetail.StartTime = spiders.StartTime
		spiderDetail.EndTime = spiders.EndTime
		spiderDetail.Project = project
	}
	return &spiderDetail, nil
}

//func readLastLinesFromLog(project, jobID string, numLines int) ([]string, error) {
//	logContent, err := app.Server.Scrapyd.FetchScrapydLog(project, jobID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to fetch log: %w", err)
//	}
//
//	lines := make([]string, 0, numLines)
//	var currentLine string
//	lineCount := 0
//	reader := bufio.NewReader(strings.NewReader(logContent))
//	for {
//		currentLine, err = reader.ReadString('\n')
//		if err != nil {
//			break
//		}
//
//		if lineCount >= numLines {
//			lines = lines[1:]
//		}
//		lines = append(lines, currentLine)
//		lineCount++
//	}
//	return lines, nil
//}

func countCrawledPagesFromLog(logContent []string) (uint64, error) {
	re := regexp.MustCompile(`\|\s*(\d+)`)

	var crawlCount uint64 = 0

	for _, line := range logContent {
		line = strings.TrimSpace(line)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			crawlCountInt, err := strconv.Atoi(matches[1])
			if err != nil {
				return crawlCount, fmt.Errorf("failed to parse crawled pages count: %w", err)
			}
			crawlCount = uint64(crawlCountInt)
		}
	}

	return crawlCount, nil
}

func countStatusCodesFromLog(logContent []string, statusCodeMap map[int]string) ([]*types.StatusCode, error) {
	re := regexp.MustCompile(`\[\w+]\s*\|\s*(\{.*})$`)
	var statusCounts []types.StatusCode

	for _, line := range logContent {
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
