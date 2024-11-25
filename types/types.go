package types

import (
	"general_spider_controll_panel/types/models"
)

type NavbarItems struct {
	Key string
	Url string
}

type ScrapydResponseGetingSpiders struct {
	NodeName string `json:"node_name"`
	Status   string `json:"status"`
	Pending  []struct {
		Id       string                 `json:"id"`
		Project  string                 `json:"project"`
		Spider   string                 `json:"spider"`
		Version  string                 `json:"version"`
		Settings map[string]interface{} `json:"settings"`
		Args     map[string]interface{} `json:"args"`
	} `json:"pending"`
	Running []struct {
		Id        string `json:"id"`
		Project   string `json:"project"`
		Spider    string `json:"spider"`
		Pid       int    `json:"pid"`
		StartTime string `json:"start_time"`
		LogUrl    string `json:"log_url"`
		ItemsUrl  string `json:"items_url"`
	} `json:"running"`
	Finished []struct {
		Id        string `json:"id"`
		Project   string `json:"project"`
		Spider    string `json:"spider"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		LogUrl    string `json:"log_url"`
		ItemsUrl  string `json:"items_url"`
	} `json:"finished"`
}

type DomainStats struct {
	Domain         string `json:"domain"`
	LastCrawled    string `json:"last_crawled"`
	ActiveSpider   uint64 `json:"active_spider"`
	PendingSpider  uint64 `json:"pending_spider"`
	FinishedSpider uint64 `json:"finished_spider"`
}

type Spider struct {
	Id        string `json:"id"`
	Project   string `json:"project"`
	Spider    string `json:"spider"`
	Pid       int    `json:"pid"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	LogUrl    string `json:"log_url"`
	ItemsUrl  string `json:"items_url"`
	Status    string `json:"status"`
	NodeName  string `json:"node_name"`
}

type SpiderDetail struct {
	Cpu           string `json:"cpu"`
	Mem           uint64 `json:"mem"`
	NodeName      string `json:"node_name"`
	PID           int    `json:"pid"`
	CrawledCount  uint64 `json:"crawled_count"`
	CrawledDetail []StatusCode
	Log           []string `json:"log"`
	Status        string   `json:"status"`
	Name          string   `json:"name"`
	Id            string   `json:"id"`
	StartTime     string   `json:"start_time"`
	EndTime       string   `json:"end_time"`
}

type StatusCode struct {
	Code      uint
	Detail    string
	Count     uint
	BaseGroup string
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Config struct {
	BaseURL string   `json:"base_url"`
	Configs []string `json:"configs"`
}

type ScrapydResponse struct {
	Jobid string `json:"jobid"`
}

type ConfigDetail struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type JobIDDetail struct {
	JobID string `json:"jobid"`
	StatusCode
	CrawledCount uint64 `json:"crawled_count"`
}

type ProxyStatus string

const (
	Online    ProxyStatus = "Online"
	Offline   ProxyStatus = "Offline"
	Checking  ProxyStatus = "Checking"
	Unchecked ProxyStatus = "Unchecked"
)

type Proxy struct {
	Address  string      `json:"address"`
	Port     string      `json:"port"`
	Protocol string      `json:"protocol"`
	Status   ProxyStatus `json:"status"`
}

type Database interface {
	CreateConfig(config *models.Config) error
	GetConfigByID(id string) (*models.Config, error)
	GetConfigsIDByDomain(domain string) ([]string, error)
	GetConfigNameAndIDByDomain(domain string) ([]*ConfigDetail, error)
	GetDomains() ([]string, error)
	GetConfigs() ([]*models.Config, error)

	GetProxies() ([]*models.Proxy, error)
	GetProxyByID(id string) (*models.Proxy, error)
	CreateProxy(proxy *models.Proxy) (*models.Proxy, error)
	UpdateProxyStatus(addr string, status models.ProxyStatus) error
	RemoveProxy(addr string) error
}
