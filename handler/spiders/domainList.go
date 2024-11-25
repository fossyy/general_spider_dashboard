package handlerSpidersDomainList

import (
	"fmt"
	"general_spider_controll_panel/app"
	HandlerSpiders "general_spider_controll_panel/handler/spiders/spider"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	domainListView "general_spider_controll_panel/view/spider"
	"net/http"
	"time"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-domains":
			domains, err := app.Server.Database.GetDomains()
			domainsDetail := GetDomainsStats(domains)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			domainListView.GetDomains(domainsDetail).Render(r.Context(), w)
			return
		}
	}
	domainListView.Main("Spiders page").Render(r.Context(), w)
}

func GetDomainsStats(domains []string) []*types.DomainStats {
	var data []*types.DomainStats
	layout := "2006-01-02 15:04:05.999999"
	if len(domains) == 0 {
		return data
	}
	for _, domain := range domains {
		spiders, err := HandlerSpiders.GetRunningSpiders(scrapydURL, domain)
		if err != nil {
			data = append(data, &types.DomainStats{
				Domain:         domain,
				LastCrawled:    "Error geting the data",
				ActiveSpider:   0,
				PendingSpider:  0,
				FinishedSpider: 0,
			})
			fmt.Println(err.Error())
			continue
		}
		if len(spiders.Running) == 0 && len(spiders.Finished) == 0 {
			data = append(data, &types.DomainStats{
				Domain:         domain,
				LastCrawled:    "Never been run before :)",
				ActiveSpider:   0,
				PendingSpider:  0,
				FinishedSpider: 0,
			})
			continue
		}
		if len(spiders.Running) > 0 {
			data = append(data, &types.DomainStats{
				Domain:         domain,
				LastCrawled:    "Now",
				ActiveSpider:   uint64(len(spiders.Running)),
				PendingSpider:  uint64(len(spiders.Pending)),
				FinishedSpider: uint64(len(spiders.Finished)),
			})
		} else {
			var recent time.Time
			for _, finishedSpider := range spiders.Finished {
				parsedTime, err := time.ParseInLocation(layout, finishedSpider.EndTime, time.Local)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				if recent.UnixMilli() < parsedTime.UnixMilli() {
					recent = parsedTime
				}
			}
			data = append(data, &types.DomainStats{
				Domain:         domain,
				LastCrawled:    utils.ConvertTIme(recent.UnixMilli()),
				ActiveSpider:   0,
				PendingSpider:  uint64(len(spiders.Pending)),
				FinishedSpider: uint64(len(spiders.Finished)),
			})
		}
	}
	return data
}
