package handlerDeleteConfig

import (
	"general_spider_controll_panel/app"
	"net/http"
)

func DELETE(w http.ResponseWriter, r *http.Request) {
	configID := r.PathValue("id")
	if err := app.Server.Database.DeleteConfigByID(configID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.Server.Logger.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
