package kafkaTopicAddHandler

import (
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types/models"
	kafkaTopicAddView "general_spider_controll_panel/view/kafka/topic/add"
	"github.com/google/uuid"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	kafkaTopicAddView.Main("Kafka topic page").Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	topicName := r.FormValue("topicName")

	if topicName == "" {
		err := app.Server.Response.SendMessageToast(w, &app.BackendResponse{
			Message: "Invalid topic name is given",
			Type:    app.Error,
			Timeout: 5000,
		})
		if err != nil {
			app.Server.Logger.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	if app.Server.Database.IsTopicPresent(topicName) {
		err := app.Server.Response.SendMessageToast(w, &app.BackendResponse{
			Message: "Topic already exists",
			Type:    app.Error,
			Timeout: 5000,
		})
		if err != nil {
			app.Server.Logger.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	err = app.Server.Database.CreateKafkaTopic(&models.KafkaTopic{
		TopicID:   uuid.New(),
		TopicName: topicName,
	})
	if err != nil {
		app.Server.Logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Hx-Redirect", "/kafka/topic")
	w.WriteHeader(http.StatusCreated)
}
