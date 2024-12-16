package kafkaBrokerAddHandler

import (
	"general_spider_controll_panel/app"
	"general_spider_controll_panel/types/models"
	"general_spider_controll_panel/view/kafka/broker/add"
	"github.com/google/uuid"
	"net"
	"net/http"
	"strings"
)

func GET(w http.ResponseWriter, r *http.Request) {
	kafkaAddView.Main("Kafka config page").Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.Server.Logger.Println("Cannot parse form : ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	kafkaBrokersGroup := r.FormValue("brokersGroup")
	kafkaBrokers := r.FormValue("brokers")

	brokers := strings.Split(kafkaBrokers, ",")
	for _, broker := range brokers {
		host, port, err := net.SplitHostPort(broker)
		if err != nil {
			app.Server.Logger.Printf("Cannot parse broker url : %s", err.Error())
			err := app.Server.Response.SendMessageToast(w, &app.BackendResponse{
				Message: "Cannot parse broker url : Invalid address is given",
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
		if host == "" || port == "" {
			app.Server.Logger.Printf("Cannot parse broker url : Invalid address is given (%s:%s)", host, port)
			err := app.Server.Response.SendMessageToast(w, &app.BackendResponse{
				Message: "Cannot parse broker url : Invalid address is given",
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
		err = app.Server.Database.CreateKafkaBroker(&models.KafkaBroker{
			BrokerID:    uuid.New(),
			BrokerGroup: kafkaBrokersGroup,
			Host:        host,
			Port:        port,
		})
		if err != nil {
			app.Server.Logger.Println("Cannot create kafka broker : ", err.Error())
			err := app.Server.Response.SendMessageToast(w, &app.BackendResponse{
				Message: "Cannot create kafka broker : " + err.Error(),
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
	}
	w.Header().Set("Hx-Redirect", "/kafka/broker")
	w.WriteHeader(http.StatusCreated)
	return
}
