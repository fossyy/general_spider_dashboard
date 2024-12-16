package app

import (
	"fmt"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/types/models"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Tools struct {
	scrapyd  *ScrapydStruct
	database types.Database
	logger   *log.Logger
}

func (t *Tools) CheckProxy(proxy *models.Proxy) {
	if proxy.Usage != "" {
		var projects []string
		exists := make(map[string]bool)

		scrapydProjects, err := t.scrapyd.GetAllProjects()
		if err != nil {
			return
		}
		crons, err := t.database.GetCrons()
		if err != nil {
			return
		}

		for _, project := range scrapydProjects {
			if !exists[project] {
				exists[project] = true
				projects = append(projects, project)
			}
		}

		for _, project := range crons {
			if !exists[project.Project] {
				exists[project.Project] = true
				projects = append(projects, project.Project)
			}
		}

		proxySpider, err := t.scrapyd.GetSpider(proxy.Usage, projects)
		if err != nil {
			return
		}

		inSchedule := false
		for _, project := range crons {
			if project.JobID == proxy.Usage {
				inSchedule = true
				break
			}
		}

		if proxySpider == nil && !inSchedule {
			err := t.database.RemoveProxyUsedStatus(proxy.Address)
			if err != nil {
				return
			}
			proxy.Usage = ""
		}
		if proxySpider != nil && proxySpider.Status != "Running" && !inSchedule {
			err := t.database.RemoveProxyUsedStatus(proxy.Address)
			if err != nil {
				return
			}
			proxy.Usage = ""
		}
	}
	err := t.database.UpdateProxyStatus(proxy.Address, models.Checking)
	rawProxy := fmt.Sprintf("%s://%s:%s", proxy.Protocol, proxy.Address, proxy.Port)
	if err != nil {
		t.logger.Println("Error while checking Proxy:", err)
		return
	}
	proxyUrl, _ := url.Parse(rawProxy)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get("https://example.com")
	if err != nil {
		ignorable := []string{
			"connection refused",
			"unreachable",
			"no route to host",
			"invalid argument",
		}
		for _, ignore := range ignorable {
			if strings.Contains(err.Error(), ignore) {
				t.logger.Printf("Proxy : %s is offline", proxyUrl)
				err := t.database.UpdateProxyStatus(proxy.Address, models.Offline)
				if err != nil {
					t.logger.Println("Error while updating Proxy status : ", err)
					return
				}
				return
			}
		}
		t.logger.Println("Error while checking Proxy : ", err)
		return
	}
	defer resp.Body.Close()

	if proxy.Usage != "" {
		err = t.database.UpdateProxyStatus(proxy.Address, models.Used)
		if err != nil {
			t.logger.Println(err)
			return
		}
	} else {
		err = t.database.UpdateProxyStatus(proxy.Address, models.Online)
		if err != nil {
			t.logger.Println(err)
			return
		}
	}
}
