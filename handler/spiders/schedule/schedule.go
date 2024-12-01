package scheduleDetailsHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/utils"
	scheduleView "general_spider_controll_panel/view/spider/schedule"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GET(w http.ResponseWriter, r *http.Request) {
	project := r.PathValue("project")
	id := r.PathValue("id")
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-schedule-action":
			scheduleView.ScheduleActionUI(id).Render(r.Context(), w)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	var jobData *types.Cron
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
			if tag == project && job.ID().String() == id {
				dbCron, err := app.Server.Database.GetCronByID(job.ID().String())
				if err != nil {
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

				jobData = &types.Cron{
					CreatedAt:      dbCron.CreatedAt.Format("Monday, 02 January 2006"),
					UpdatedAt:      dbCron.UpdatedAt.Format("Monday, 02 January 2006"),
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
				}
			}
		}
	}
	if jobData == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	timelines, err := app.Server.Database.GetTimelineByContext(jobData.JobId)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	scheduleView.Main("schedule spider page", jobData, timelines).Render(r.Context(), w)
	return
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	project := r.PathValue("project")
	id := r.PathValue("id")
	cron, err := app.Server.Database.GetCronByID(id)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = app.Server.Cron.RemoveJob(cron.ID)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = app.Server.Database.RemoveScheduleByID(id)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = app.Server.Database.RemoveTimelineByContext(cron.JobID)
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = app.Server.Database.RemoveProxyUsedStatusByJobID(cron.JobID)
	if err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			w.Header().Set("HX-Redirect", fmt.Sprintf("/spiders/%s?type=schedule", project))
			return
		}
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", fmt.Sprintf("/spiders/%s?type=schedule", cron.Project))
	w.WriteHeader(http.StatusOK)
	return
}
