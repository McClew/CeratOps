package controller

import (
	"CeratOps/view"
	"net/http"
)

func DashboardHandler(writer http.ResponseWriter, request *http.Request) {
	view.DashboardPage().Render(request.Context(), writer)
}
