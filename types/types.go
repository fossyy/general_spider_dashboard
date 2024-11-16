package types

import "general_spider_controll_panel/types/models"

type NavbarItems struct {
	Key string
	Url string
}

type ScrapydResponseGetingSpiders struct {
	Pending []interface{} `json:"pending"`
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
	Status   string `json:"status"`
	NodeName string `json:"node_name"`
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
}

type SpiderDetails struct {
	NodeName  string `json:"node_name"`
	Status    string `json:"status"`
	Currstate string `json:"currstate"`
	Spider    *Spider
	Detail    struct {
		Cpu           string `json:"cpu"`
		Mem           uint64 `json:"mem"`
		NodeName      string `json:"node_name"`
		CrawledCount  uint64 `json:"crawled_count"`
		CrawledDetail []StatusCode
	}
	Log []string `json:"log"`
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

type Database interface {
	CreateConfig(config *models.Config) error
	GetConfigByID(id uint) (*models.Config, error)
}
