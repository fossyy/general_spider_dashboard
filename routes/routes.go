package routes

import (
	handlerPreview "general_spider_controll_panel/handler/api/preview"
	configHandler "general_spider_controll_panel/handler/config"
	handlerConfigByID "general_spider_controll_panel/handler/config/configByID"
	handlerConfigs "general_spider_controll_panel/handler/configs"
	deployHandler "general_spider_controll_panel/handler/deploy"
	"general_spider_controll_panel/handler/kafka/broker"
	"general_spider_controll_panel/handler/kafka/broker/add"
	kafkaTopicHandler "general_spider_controll_panel/handler/kafka/topic"
	kafkaTopicAddHandler "general_spider_controll_panel/handler/kafka/topic/add"
	proxiesHandler "general_spider_controll_panel/handler/proxies"
	handlerSpidersDomainList "general_spider_controll_panel/handler/spiders"
	handlerSpiderDetails "general_spider_controll_panel/handler/spiders/details"
	scheduleDetailsHandler "general_spider_controll_panel/handler/spiders/schedule"
	HandlerSpiders "general_spider_controll_panel/handler/spiders/spider"
	"net/http"
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

	kafkaRouter := http.NewServeMux()
	handler.Handle("/kafka/", http.StripPrefix("/kafka", kafkaRouter))
	kafkaRouter.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/kafka/broker", http.StatusSeeOther)
	})
	kafkaRouter.HandleFunc("GET /broker", broker.GET)
	kafkaRouter.HandleFunc("GET /broker/add", kafkaBrokerAddHandler.GET)
	kafkaRouter.HandleFunc("POST /broker/add", kafkaBrokerAddHandler.POST)
	kafkaRouter.HandleFunc("GET /topic", kafkaTopicHandler.GET)
	kafkaRouter.HandleFunc("GET /topic/add", kafkaTopicAddHandler.GET)
	kafkaRouter.HandleFunc("POST /topic/add", kafkaTopicAddHandler.POST)

	handler.Handle("/results/", http.StripPrefix("/results/", http.FileServer(http.Dir("./results/"))))
	handler.Handle("/logs/", http.StripPrefix("/logs/", http.FileServer(http.Dir("./logs/"))))

	handler.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return handler
}
