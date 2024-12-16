package kafkaTopicHandler

import (
	"general_spider_controll_panel/app"
	kafkaTopicView "general_spider_controll_panel/view/kafka/topic"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	topics, err := app.Server.Database.GetKafkaTopics()
	if err != nil {
		app.Server.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	kafkaTopicView.Main("Kafka topic page", topics).Render(r.Context(), w)
}
