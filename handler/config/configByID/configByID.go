package handlerConfigByID

import (
	"encoding/json"
	"general_spider_controll_panel/app"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config, err := app.Server.Database.GetConfigByID(id)
	if err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(config); err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	return
}
