package proxiesHandler

import (
	"fmt"
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types/models"
	proxiesView "general_spider_controll_panel/view/proxies"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
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
					CheckProxy(proxy)
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
			go CheckProxy(proxy)
			proxy.Status = models.Checking
			proxiesView.Proxy(proxy).Render(r.Context(), w)
			return
		default:
			fmt.Println("anjir")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	r.ParseForm()
	proxyAddress := r.Form.Get("proxyAddr")
	proxyProto := r.Form.Get("proxyProto")

	spliteAddr := strings.Split(proxyAddress, ":")
	proxy := &models.Proxy{
		Address:  spliteAddr[0],
		Port:     spliteAddr[1],
		Protocol: proxyProto,
	}
	proxy, err := app.Server.Database.CreateProxy(proxy)
	if err != nil {
		fmt.Println(err)
		fmt.Println(spliteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go CheckProxy(proxy)
	proxiesView.Proxy(proxy).Render(r.Context(), w)
	return
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

func CheckProxy(proxy *models.Proxy) {
	err := app.Server.Database.UpdateProxyStatus(proxy.Address, models.Checking)
	rawProxy := fmt.Sprintf("%s://%s:%s", proxy.Protocol, proxy.Address, proxy.Port)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while checking Proxy:", err)
		return
	}
	proxyUrl, _ := url.Parse(rawProxy)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get("https://example.com")
	if err != nil || resp.StatusCode != http.StatusOK {
		err := app.Server.Database.UpdateProxyStatus(proxy.Address, models.Offline)
		if err != nil {
			fmt.Println("Error while checking Proxy 2:", err)
			return
		}
		fmt.Println("Error while checking Proxy 3:", err)
		return
	}
	defer resp.Body.Close()

	err = app.Server.Database.UpdateProxyStatus(proxy.Address, models.Online)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Proxy %s:%s is online.\n", proxy.Address, &proxy.Port)
}
