package downloadHandler

import (
	"general_spider_controll_panel/app"
	downloadView "general_spider_controll_panel/view/download"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	datas, err := app.Server.Scrapyd.GetLogsAndResults()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		app.Server.Logger.Println(err.Error())
		return
	}

	downloadView.Main("Download Page", datas).Render(r.Context(), w)
}
