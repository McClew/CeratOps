package controller

import (
	view "CeratOps/view/dashboard"
	"net/http"
)

func DashboardHandler(writer http.ResponseWriter, request *http.Request) {
	view.DashboardPage().Render(request.Context(), writer)
}
