package app

import (
	_ "embed"
	"encoding/json"
	"general_spider_controll_panel/types"
	"general_spider_controll_panel/types/models"
	"general_spider_controll_panel/utils"
	"github.com/go-co-op/gocron/v2"
	"log"
	"net/http"
	"sync"
	"time"
)

type BackendResponseType string

const (
	Error   BackendResponseType = "error"
	Success BackendResponseType = "success"
	Info    BackendResponseType = "info"
)

type BackendResponse struct {
	Message string              `json:"message"`
	Type    BackendResponseType `json:"type"`
	Timeout uint                `json:"timeout"`
}

var Server *App

//go:embed general.egg
var egg []byte

type App struct {
	http.Server
	Database types.Database
	Logger   *log.Logger
	Scrapyd  *ScrapydStruct
	Tools    *Tools
	Cron     gocron.Scheduler
	Response BackendResponse
}

func NewApp(addr string, handler http.Handler, database types.Database, cron gocron.Scheduler, logger *log.Logger) *App {
	scrapyd := ScrapydStruct{
		ScrapydURL: utils.Getenv("SCRAPYD_URL"),
		Version:    "1.0",
		Spider:     "general_engine",
	}

	tools := Tools{
		scrapyd:  &scrapyd,
		database: database,
		logger:   logger,
	}

	go func() {
		if err := scrapyd.UploadEggToAllProjects(database); err != nil {
			logger.Fatal("Error: " + err.Error())
		}
	}()

	go cron.Start()

	go func() {
		crons, err := database.GetCrons()
		if err != nil {
			logger.Fatal("Error: " + err.Error())
			return
		}

		for _, dbCron := range crons {
			var additionalArgs map[string]string
			err := json.Unmarshal(dbCron.AdditionalArgs, &additionalArgs)
			if err != nil {
				logger.Fatal("Error: " + err.Error())
				return
			}
			jobTask := gocron.NewTask(func(project string, additionalArgs map[string]string) {
				scrapydSpider, err := scrapyd.GetSpider(dbCron.JobID, []string{project})
				if err != nil {
					err := database.CreateTimeline(&models.Timeline{
						Title:   "Job Failed",
						Message: "Skipping deployment, Error while deploying, Error : " + err.Error(),
						Context: dbCron.JobID,
						Status:  models.Failed,
					})
					if err != nil {
						logger.Println("Error creating timeline: " + err.Error())
						return
					}
					logger.Println(err)
					return
				}
				if scrapydSpider == nil || scrapydSpider.Status != "Running" {
					_, err = scrapyd.RunSpider(project, additionalArgs)
					if err != nil {
						err := database.CreateTimeline(&models.Timeline{
							Title:   "Job Failed",
							Message: "Skipping deployment, Error while deploying, Error : " + err.Error(),
							Context: dbCron.JobID,
							Status:  models.Failed,
						})
						if err != nil {
							logger.Println("Error creating timeline: " + err.Error())
							return
						}
						logger.Println(err)
						return
					}
					err := database.CreateTimeline(&models.Timeline{
						Title:   "Job Started",
						Message: "Cron job execution completed successfully",
						Context: dbCron.JobID,
						Status:  models.Success,
					})
					if err != nil {
						logger.Println("Error creating timeline: " + err.Error())
						return
					}
				} else {
					err := database.CreateTimeline(&models.Timeline{
						Title:   "Job Failed",
						Message: "Skipping deployment, last deployment still running",
						Context: dbCron.JobID,
						Status:  models.Failed,
					})
					if err != nil {
						logger.Println("Error creating timeline: " + err.Error())
						return
					}
					return
				}
			}, dbCron.Project, additionalArgs)
			jobType := gocron.CronJob(dbCron.Schedule, false)
			job, err := cron.NewJob(jobType, jobTask, gocron.WithTags(dbCron.Project))
			logger.Println("Adding cron from db to stack : ", dbCron.ID, dbCron.AdditionalArgs)
			if err != nil {
				logger.Fatal("Error: " + err.Error())
				return
			}
			err = database.ChangeCronID(dbCron.ID.String(), job.ID())
			if err != nil {
				logger.Fatal("Error: " + err.Error())
				return
			}
		}
	}()

	go func() {
		for {
			proxies, err := database.GetProxies()
			if err != nil {
				logger.Fatal("Error: " + err.Error())
				return
			}
			wg := sync.WaitGroup{}
			for _, proxy := range proxies {
				wg.Add(1)
				go func() {
					defer wg.Done()
					logger.Println("Checking proxy : ", proxy.Address)
					tools.CheckProxy(proxy)
				}()
			}
			wg.Wait()
			time.Sleep(5 * time.Minute)
		}
	}()

	return &App{
		Server: http.Server{
			Addr:    addr,
			Handler: handler,
		},
		Database: database,
		Cron:     cron,
		Logger:   logger,
		Scrapyd:  &scrapyd,
		Tools:    &tools,
		Response: BackendResponse{},
	}
}

func (be *BackendResponse) SendMessageToast(w http.ResponseWriter, message *BackendResponse) error {
	if message.Timeout < 1000 {
		message.Timeout = 1000
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		return err
	}
	return nil
}
