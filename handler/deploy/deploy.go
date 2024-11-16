package deployHandler

import (
	deployView "general_spider_controll_panel/view/deploy"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	deployView.Main("Deploy page").Render(r.Context(), w)
}
