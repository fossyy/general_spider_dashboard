package proxiesHandler

import (
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types/models"
	proxiesView "general_spider_controll_panel/view/proxies"
	"net/http"
	"strings"
	"sync"
)

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "get-proxies":
			proxies, err := app.Server.Database.GetProxies()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			proxiesView.Proxies(proxies).Render(r.Context(), w)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	proxiesView.Main("Proxies Page").Render(r.Context(), w)
	return
}

func POST(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("hx-request") == "true" {
		switch r.URL.Query().Get("action") {
		case "test-proxies":
			wg := &sync.WaitGroup{}
			proxies, err := app.Server.Database.GetProxies()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			for _, proxy := range proxies {
				proxy.Status = models.Checking
				wg.Add(1)
				go func(proxy *models.Proxy) {
					defer wg.Done()
					app.Server.Tools.CheckProxy(proxy)
				}(proxy)
			}

			wg.Wait()
			proxiesView.Proxies(proxies).Render(r.Context(), w)
			return
		case "test-proxy":
			id := r.URL.Query().Get("id")
			proxy, err := app.Server.Database.GetProxyByID(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			go app.Server.Tools.CheckProxy(proxy)
			proxy.Status = models.Checking
			proxiesView.Proxy(proxy).Render(r.Context(), w)
			return
		case "add-proxy":
			r.ParseForm()
			proxyAddress := r.Form.Get("proxyAddr")
			proxyProto := r.Form.Get("proxyProto")

			splitAddr := strings.Split(proxyAddress, ":")
			proxy := &models.Proxy{
				Address:  splitAddr[0],
				Port:     splitAddr[1],
				Protocol: proxyProto,
			}
			proxy, err := app.Server.Database.CreateProxy(proxy)
			if err != nil {
				app.Server.Logger.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			go app.Server.Tools.CheckProxy(proxy)
			proxiesView.Proxy(proxy).Render(r.Context(), w)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	addr := r.PathValue("addr")
	err := app.Server.Database.RemoveProxy(addr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
