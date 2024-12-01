package handlerPreview

import (
	"fmt"
	configHandler "general_spider_controll_panel/handler/config"
	"io"
	"net/http"
)

var data map[string]interface{}

func init() {
	data = make(map[string]interface{})
}

func GET(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	dataBytes, ok := data[id].([]byte)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataBytes)
	return
}

func POST(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body err: %v", err)
		return
	}

	configHandler.TestRun[id] = body
	fmt.Fprintf(w, string(body))
}
