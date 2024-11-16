package routes

import (
	configHandler "general_spider_controll_panel/handler/config"
	handlerConfigByID "general_spider_controll_panel/handler/config/configByID"
	deployHandler "general_spider_controll_panel/handler/deploy"
	handlerSpider "general_spider_controll_panel/handler/spider"
	handlerSpiders "general_spider_controll_panel/handler/spiders"
	handlerSpiderDetails "general_spider_controll_panel/handler/spiders/details"
	"net/http"
)

func Setup() *http.ServeMux {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /config", configHandler.GET)
	handler.HandleFunc("POST /config", configHandler.POST)
	handler.HandleFunc("GET /config/{id}", handlerConfigByID.GET)
	handler.HandleFunc("GET /spiders", handlerSpiders.GET)
	handler.HandleFunc("POST /spider/{config}", handlerSpider.POST)
	handler.HandleFunc("GET /spider/{id}", handlerSpiderDetails.GET)
	handler.HandleFunc("GET /deploy", deployHandler.GET)
	handler.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return handler
}
