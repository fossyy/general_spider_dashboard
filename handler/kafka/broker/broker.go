package broker

import (
	"general_spider_controll_panel/app"
	kafkaView "general_spider_controll_panel/view/kafka/broker"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	brokers, err := app.Server.Database.GetKafkaBrokers()
	if err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	kafkaView.Main("Kafka config page", brokers).Render(r.Context(), w)
}
