package kafkaBrokerDeleteHandler

import (
	"general_spider_controll_panel/app"
	"net/http"
)

func DELETE(w http.ResponseWriter, r *http.Request) {
	kafkaID := r.PathValue("id")
	if err := app.Server.Database.DeleteKafkaBroker(kafkaID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		app.Server.Logger.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
