package routes

import (
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	handlerPreview "general_spider_controll_panel/handler/api/preview"
	configHandler "general_spider_controll_panel/handler/config"
	handlerConfigByID "general_spider_controll_panel/handler/config/configByID"
	handlerConfigs "general_spider_controll_panel/handler/configs"
	deployHandler "general_spider_controll_panel/handler/deploy"
	proxiesHandler "general_spider_controll_panel/handler/proxies"
	handlerSpidersDomainList "general_spider_controll_panel/handler/spiders"
	handlerSpiderDetails "general_spider_controll_panel/handler/spiders/details"
	scheduleDetailsHandler "general_spider_controll_panel/handler/spiders/schedule"
	HandlerSpiders "general_spider_controll_panel/handler/spiders/spider"
	"net/http"
	"time"
)

func Setup() *http.ServeMux {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /config", configHandler.GET)
	handler.HandleFunc("POST /config", configHandler.POST)
	handler.HandleFunc("GET /config/{id}", handlerConfigByID.GET)
	handler.HandleFunc("GET /configs", handlerConfigs.GET)
	handler.HandleFunc("GET /spiders", handlerSpidersDomainList.GET)
	handler.HandleFunc("GET /spiders/{project}", HandlerSpiders.GET)
	handler.HandleFunc("GET /spiders/{project}/active/{id}", handlerSpiderDetails.GET)
	handler.HandleFunc("GET /spiders/{project}/schedule/{id}", scheduleDetailsHandler.GET)
	handler.HandleFunc("DELETE /spiders/{project}/schedule/{id}", scheduleDetailsHandler.DELETE)
	handler.HandleFunc("GET /proxies", proxiesHandler.GET)
	handler.HandleFunc("POST /proxies", proxiesHandler.POST)
	handler.HandleFunc("DELETE /proxies/{addr}", proxiesHandler.DELETE)
	handler.HandleFunc("DELETE /spiders/{project}/active/{id}", handlerSpiderDetails.DELETE)
	handler.HandleFunc("GET /api/preview/{id}", handlerPreview.GET)
	handler.HandleFunc("POST /api/preview/{id}", handlerPreview.POST)
	handler.HandleFunc("GET /deploy", deployHandler.GET)
	handler.HandleFunc("POST /deploy", deployHandler.POST)
	handler.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		crons, err := app.Server.Database.GetCrons()
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(crons)
	})
	handler.HandleFunc("GET /test2", func(w http.ResponseWriter, r *http.Request) {
		var jobData []map[string]interface{}
		for _, job := range app.Server.Cron.Jobs() {
			run, err := job.NextRun()
			if err != nil {
				return
			}
			lastRun, err := job.LastRun()
			if err != nil {
				return
			}
			jobData = append(jobData, map[string]interface{}{
				"ID":        job.ID().String(),
				"Name":      job.Name(),
				"LastRun":   lastRun,
				"NextRun":   run,
				"Countdown": run.Sub(time.Now()).Seconds(),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(jobData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	handler.HandleFunc("GET /test3", func(w http.ResponseWriter, r *http.Request) {
		err := app.Server.Database.RemoveTimelineByContext("local_eg3gjyrrk751z2m3dekwpuulzr8j2")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})
	handler.Handle("/results/", http.StripPrefix("/results/", http.FileServer(http.Dir("./results/"))))
	handler.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("./logs/"))))

	handler.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return handler
}
