package handlerSpidersDomainList

import (
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	domainListView "general_spider_controll_panel/view/spider"
	"net/http"
	"sort"
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
				app.Server.Logger.Println(err.Error())
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
		spiders, err := app.Server.Scrapyd.GetRunningSpiders(domain)
		if err != nil {
			data = append(data, &types.DomainStats{
				Domain:          domain,
				LastCrawled:     "Error getting the data",
				ActiveSpider:    0,
				PendingSpider:   0,
				FinishedSpider:  0,
				LastCrawledAt:   time.Time{},
				ScheduledSpider: 0,
			})
			app.Server.Logger.Println(err.Error())
			continue
		}
		scheduledSpiders, err := app.Server.Database.CountScheduledSpiders(domain)
		if err != nil {
			data = append(data, &types.DomainStats{
				Domain:          domain,
				LastCrawled:     "Error getting the data",
				ActiveSpider:    0,
				PendingSpider:   0,
				FinishedSpider:  0,
				LastCrawledAt:   time.Time{},
				ScheduledSpider: 0,
			})
			app.Server.Logger.Println(err.Error())
			continue
		}
		if len(spiders.Running) == 0 && len(spiders.Finished) == 0 {
			data = append(data, &types.DomainStats{
				Domain:          domain,
				LastCrawled:     "Never been run before :)",
				ActiveSpider:    0,
				PendingSpider:   0,
				FinishedSpider:  0,
				LastCrawledAt:   time.Time{},
				ScheduledSpider: scheduledSpiders,
			})
			continue
		}
		if len(spiders.Running) > 0 {
			data = append(data, &types.DomainStats{
				Domain:          domain,
				LastCrawled:     "Now",
				ActiveSpider:    uint64(len(spiders.Running)),
				PendingSpider:   uint64(len(spiders.Pending)),
				FinishedSpider:  uint64(len(spiders.Finished)),
				LastCrawledAt:   time.Now(),
				ScheduledSpider: scheduledSpiders,
			})
		} else {
			var recent time.Time
			for _, finishedSpider := range spiders.Finished {
				parsedTime, err := time.ParseInLocation(layout, finishedSpider.EndTime, time.Local)
				if err != nil {
					app.Server.Logger.Println(err.Error())
					continue
				}
				if recent.UnixMilli() < parsedTime.UnixMilli() {
					recent = parsedTime
				}
			}
			data = append(data, &types.DomainStats{
				Domain:          domain,
				LastCrawled:     utils.ConvertTIme(recent.UnixMilli()),
				ActiveSpider:    0,
				PendingSpider:   uint64(len(spiders.Pending)),
				FinishedSpider:  uint64(len(spiders.Finished)),
				LastCrawledAt:   recent,
				ScheduledSpider: scheduledSpiders,
			})
		}
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].LastCrawledAt.After(data[j].LastCrawledAt)
	})
	return data
}
