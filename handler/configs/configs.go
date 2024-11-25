package handlerConfigs

import (
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/app"
	"net/http"
)

func GET(w http.ResponseWriter, req *http.Request) {
	configs, err := app.Server.Database.GetConfigs()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	marshal, err := json.Marshal(configs)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(marshal)
}
