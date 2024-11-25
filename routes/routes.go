package routes

import (
	handlerPreview "general_spider_controll_panel/handler/api/preview"
	configHandler "general_spider_controll_panel/handler/config"
	handlerConfigByID "general_spider_controll_panel/handler/config/configByID"
	handlerConfigs "general_spider_controll_panel/handler/configs"
	deployHandler "general_spider_controll_panel/handler/deploy"
	proxiesHandler "general_spider_controll_panel/handler/proxies"
	handlerSpidersDomainList "general_spider_controll_panel/handler/spiders"
	handlerSpiderDetails "general_spider_controll_panel/handler/spiders/details"
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
	handler.HandleFunc("GET /spiders/{project}/{id}", handlerSpiderDetails.GET)
	handler.HandleFunc("GET /proxies", proxiesHandler.GET)
	handler.HandleFunc("POST /proxies", proxiesHandler.POST)
	handler.HandleFunc("DELETE  /proxies/{addr}", proxiesHandler.DELETE)
	handler.HandleFunc("DELETE /spiders/{project}/{id}", handlerSpiderDetails.DELETE)
	handler.HandleFunc("GET /api/preview/{id}", handlerPreview.GET)
	handler.HandleFunc("POST /api/preview/{id}", handlerPreview.POST)
	handler.HandleFunc("GET /deploy", deployHandler.GET)
	handler.HandleFunc("POST /deploy", deployHandler.POST)
	handler.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	return handler
}
