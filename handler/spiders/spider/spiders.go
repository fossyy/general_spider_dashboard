package HandlerSpiders

import (
	"encoding/json"
	"errors"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	spiderView "general_spider_controll_panel/view/spider/spider"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GET(w http.ResponseWriter, r *http.Request) {
	project := r.PathValue("project")
	err := ensureProjectAndUpload(project)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-spiders":
			spiders, err := app.Server.Scrapyd.GetRunningSpiders(project)
			if err != nil {
				return
			}
			spiderView.GetSpider(spiders).Render(r.Context(), w)
			return
		case "get-scheduled":
			var jobData []*types.Cron
			for _, job := range app.Server.Cron.Jobs() {
				run, err := job.NextRun()
				if err != nil {
					return
				}
				lastRun, err := job.LastRun()
				if err != nil {
					return
				}
				for _, tag := range job.Tags() {
					if tag == project {
						dbCron, err := app.Server.Database.GetCronByID(job.ID().String())
						if err != nil {
							if errors.Is(err, gorm.ErrRecordNotFound) {
								break
							}

							app.Server.Logger.Println(err.Error())
							return
						}
						var additionalArgs map[string]string
						err = json.Unmarshal(dbCron.AdditionalArgs, &additionalArgs)
						if err != nil {
							app.Server.Logger.Println(err.Error())
							return
						}
						var proxyAddress []string
						err = json.Unmarshal(dbCron.ProxyAddresses, &proxyAddress)
						if err != nil {
							app.Server.Logger.Println(err.Error())
							return
						}

						jobData = append(jobData, &types.Cron{
							CreatedAt:      dbCron.CreatedAt.Format("15:04:05 on Monday, 02 January 2006"),
							UpdatedAt:      dbCron.UpdatedAt.Format("15:04:05 on Monday, 02 January 2006"),
							Countdown:      utils.TimeUntil(run.Sub(time.Now().In(time.Local)).Milliseconds()),
							ID:             job.ID().String(),
							LastRun:        lastRun.Format("15:04:05 on Monday, 02 January 2006"),
							Name:           job.Name(),
							NextRun:        run.Format("15:04:05 on Monday, 02 January 2006"),
							Schedule:       dbCron.Schedule,
							Project:        dbCron.Project,
							Spider:         dbCron.Spider,
							ConfigId:       dbCron.ConfigID,
							OutputDst:      dbCron.OutputDST,
							JobId:          dbCron.JobID,
							AdditionalArgs: additionalArgs,
							ProxyAddresses: proxyAddress,
							Proxies:        dbCron.Proxies,
						})
					}
				}
			}
			spiderView.GetScheduled(jobData).Render(r.Context(), w)
			return
		case "change-to-scheduled":
			spiderView.ChangeToScheduleTable().Render(r.Context(), w)
			return
		case "change-to-running":
			spiderView.ChangeToRunningTable().Render(r.Context(), w)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	spiderView.Main("Spider Page", r.URL.Query().Get("type")).Render(r.Context(), w)
	return
}

func ensureProjectAndUpload(project string) error {
	dbProjects, err := app.Server.Database.GetDomains()
	if err != nil {
		return fmt.Errorf("failed to get projects: %w", err)
	}
	scrapydProjects, err := app.Server.Scrapyd.GetAllProjects()
	if err != nil {
		return fmt.Errorf("failed to get projects: %w", err)
	}
	var projects []string
	exists := make(map[string]bool)

	for _, item := range dbProjects {
		if !exists[item] {
			exists[item] = true
			projects = append(projects, item)
		}
	}

	for _, item := range scrapydProjects {
		if !exists[item] {
			exists[item] = true
			projects = append(projects, item)
		}
	}

	projectExists := false
	for _, existingProject := range projects {
		if existingProject == project {
			projectExists = true
			break
		}
	}
	if !projectExists {
		app.Server.Logger.Printf("Project '%s' does not exist. Creating it by uploading the egg.\n", project)
		return app.Server.Scrapyd.UploadEgg(project)
	}

	return nil
}
