package controller

import (
	"CeratOps/view"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	view.DashboardPage().Render(r.Context(), w)
}
