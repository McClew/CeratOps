package controller

import (
	view "CeratOps/view/incidents"
	"net/http"
)

func IncidentsHandler(writer http.ResponseWriter, request *http.Request) {
	view.IncidentsPage().Render(request.Context(), writer)
}
