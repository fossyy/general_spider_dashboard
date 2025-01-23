package downloadLogHandler

import (
	"fmt"
	"general_spider_controll_panel/app"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	project := r.PathValue("project")
	jobID := r.PathValue("job_id")

	log, err := app.Server.Scrapyd.GetLog(project, jobID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		app.Server.Logger.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"combined_%s.log\"", jobID))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(log)))
	w.Write(log)
	return
}
