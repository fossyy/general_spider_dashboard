package deployHandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types/models"
	"general_spider_controll_panel/utils"
	deployView "general_spider_controll_panel/view/deploy"
	"github.com/go-co-op/gocron/v2"
	"net/http"
	"net/url"
	"strings"
)

var scrapydURL = utils.Getenv("SCRAPYD_URL")
var spider = "general_engine"
var version = "1.0"

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "query-broker":
			query := r.URL.Query().Get("q")
			brokers, err := app.Server.Database.GetKafkaBrokersByName(query)
			fmt.Println(brokers)
			if err != nil {
				app.Server.Logger.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			deployView.BrokersUI(brokers).Render(r.Context(), w)
			return
		case "get-proxies":
			proxies, err := app.Server.Database.GetActiveProxies()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			project, err := app.Server.Scrapyd.GetAllProjects()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			var finalProxy []*models.Proxy
			for _, proxy := range proxies {
				if proxy.Usage != "" {
					proxySpider, err := app.Server.Scrapyd.GetSpider(proxy.Usage, project)
					if err != nil {
						continue
					}
					if proxySpider == nil {
						finalProxy = append(finalProxy, proxy)
					}
					if proxySpider != nil && proxySpider.Status != "Running" {
						finalProxy = append(finalProxy, proxy)
					}
				} else {
					finalProxy = append(finalProxy, proxy)
				}
			}
			deployView.ProxiesUI(finalProxy).Render(r.Context(), w)
			return
		case "get-configs":
			domain := r.URL.Query().Get("domainSelect")
			configs, err := app.Server.Database.GetConfigNameAndIDByDomain(domain)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			deployView.ConfigListUI(configs).Render(r.Context(), w)
			return
		case "get-additional-output-settings":
			outputDest := r.URL.Query().Get("outputDestination")
			if outputDest == "kafka" {
				brokers, err := app.Server.Database.GetKafkaBrokers()
				if err != nil {
					app.Server.Logger.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				deployView.KafkaSettingsUI(brokers).Render(r.Context(), w)
				return
			} else {
				w.Write([]byte(""))
				return
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	domains, err := app.Server.Database.GetDomains()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	deployView.Main("Deploy page", domains).Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	configID := r.Form.Get("configSelect")
	outputDest := r.Form.Get("outputDestination")
	selectedBrokers := r.Form.Get("selectedBrokers")
	kafkaTopic := r.Form.Get("kafkaTopic")
	proxies := r.Form.Get("selectedProxies")
	cookies := r.Form.Get("cookies")

	proxyList := strings.Split(proxies, ",")
	for i, proxy := range proxyList {
		proxyList[i] = strings.ReplaceAll(proxy, " ", "")
	}
	proxiesJSON, _ := json.Marshal(proxyList)

	jobid := fmt.Sprintf("%s_%s", outputDest, utils.GenerateRandomString(32))
	for _, proxy := range proxyList {
		parsed, _ := url.Parse(proxy)

		err := app.Server.Database.UpdateProxyAsUsed(strings.Split(parsed.Host, ":")[0], jobid)
		if err != nil {
			app.Server.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	additionalArgs := map[string]string{
		"config_id":  configID,
		"output_dst": outputDest,
		"proxies":    base64.StdEncoding.EncodeToString(proxiesJSON),
		"jobid":      jobid,
	}

	if cookies != "" {
		var cookiesJson map[string]interface{}
		err := json.Unmarshal([]byte(cookies), &cookiesJson)
		if err != nil {
			app.Server.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cookiesJsonBytes, err := json.Marshal(cookiesJson)
		if err != nil {
			app.Server.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		additionalArgs["cookies"] = base64.StdEncoding.EncodeToString(cookiesJsonBytes)
	}

	if selectedBrokers != "" {
		brokersList := strings.Split(selectedBrokers, ",")
		var brokers []string
		brokerWriter := &bytes.Buffer{}
		for _, brokerID := range brokersList {
			broker, err := app.Server.Database.GetKafkaBrokersById(brokerID)
			if err != nil {
				app.Server.Logger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Println("broker : ", broker)
			brokers = append(brokers, fmt.Sprintf("%s:%s", broker.Host, broker.Port))
		}
		json.NewEncoder(brokerWriter).Encode(brokers)
		additionalArgs["kafka_server"] = base64.StdEncoding.EncodeToString(brokerWriter.Bytes())
		fmt.Println("nih : ", additionalArgs["kafka_server"])
	}
	if kafkaTopic != "" {
		additionalArgs["kafka_topic"] = kafkaTopic
	}
	project := r.Form.Get("domainSelect")

	switch r.Form.Get("deploymentTime") {
	case "now":
		_, err = app.Server.Scrapyd.RunSpider(project, additionalArgs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Hx-Redirect", fmt.Sprintf("/spiders/%s", project))
		w.WriteHeader(http.StatusCreated)
	default:
		cronExpression := r.Form.Get("cronExpression")
		if cronExpression == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jobTask := gocron.NewTask(func(project string, additionalArgs map[string]string) {
			scrapydSpider, err := app.Server.Scrapyd.GetSpider(jobid, []string{project})
			if err != nil {
				err := app.Server.Database.CreateTimeline(&models.Timeline{
					Title:   "Job Failed",
					Message: "Skipping deployment, Error while deploying, Error : " + err.Error(),
					Context: jobid,
					Status:  models.Failed,
				})
				if err != nil {
					app.Server.Logger.Println("Error creating timeline: " + err.Error())
					return
				}
				return
			}
			if scrapydSpider == nil || scrapydSpider.Status != "Running" {
				_, err = app.Server.Scrapyd.RunSpider(project, additionalArgs)
				if err != nil {
					err := app.Server.Database.CreateTimeline(&models.Timeline{
						Title:   "Job Failed",
						Message: "Skipping deployment, Error while deploying, Error : " + err.Error(),
						Context: jobid,
						Status:  models.Failed,
					})
					if err != nil {
						app.Server.Logger.Println("Error creating timeline: " + err.Error())
						return
					}
					return
				}
				err := app.Server.Database.CreateTimeline(&models.Timeline{
					Title:   "Job Started",
					Message: "Cron job execution completed successfully",
					Context: jobid,
					Status:  models.Success,
				})
				if err != nil {
					app.Server.Logger.Println("Error creating timeline: " + err.Error())
					return
				}
				return
			} else {
				err := app.Server.Database.CreateTimeline(&models.Timeline{
					Title:   "Job Failed",
					Message: "Skipping deployment, last deployment still running",
					Context: jobid,
					Status:  models.Failed,
				})
				if err != nil {
					app.Server.Logger.Println("Error creating timeline: " + err.Error())
					return
				}
				return
			}
		}, project, additionalArgs)
		jobType := gocron.CronJob(cronExpression, false)
		job, err := app.Server.Cron.NewJob(jobType, jobTask, gocron.WithTags(project))
		if err != nil {
			app.Server.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		additionalArgsMarshal, err := json.Marshal(additionalArgs)
		if err != nil {
			app.Server.Logger.Println(err)
			return
		}
		err = app.Server.Database.CreateCron(&models.Schedule{
			ID:             job.ID(),
			Schedule:       cronExpression,
			Project:        project,
			Spider:         spider,
			ConfigID:       configID,
			OutputDST:      outputDest,
			ProxyAddresses: proxiesJSON,
			JobID:          jobid,
			AdditionalArgs: additionalArgsMarshal,
		})
		if err != nil {
			app.Server.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Hx-Redirect", fmt.Sprintf("/spiders/%s?type=schedule", project))
		w.WriteHeader(http.StatusCreated)
	}
}
