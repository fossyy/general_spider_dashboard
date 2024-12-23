package handlerGetConfigVersion

import (
	"encoding/json"
	"general_spider_controll_panel/app"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	version, err := app.Server.Database.GetCombinedVersion(id)
	if err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(version)
	if err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	return
}
