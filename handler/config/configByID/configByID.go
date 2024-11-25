package handlerConfigByID

import (
	"fmt"
	"general_spider_controll_panel/app"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	config, err := app.Server.Database.GetConfigByID(id)
	if err != nil {
		fmt.Fprintf(w, "Error when geting config : %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(config.Data)
}
