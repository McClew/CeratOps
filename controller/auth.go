package controller

import (
	"CeratOps/model"
	view "CeratOps/view/auth"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	// If GET show the page
	if request.Method == http.MethodGet {
		if model.IsAuthenticated(request) {
			http.Redirect(writer, request, "/dashboard", http.StatusSeeOther)
			return
		}
		templ.Handler(view.LoginPage("")).ServeHTTP(writer, request)
		return
	}

	// If POST process the login
	email := request.FormValue("email")
	pass := request.FormValue("password")
	user := model.GetUser()

	// Check email and use internal model utility for password verification
	if email == user.Email && model.CheckPasswordHash(pass, user.PasswordHash) {
		token, err := model.GenerateJWT(email)
		if err != nil {
			templ.Handler(view.LoginPage("Error generating token")).ServeHTTP(writer, request)
			return
		}

		http.SetCookie(writer, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true, // JS can't read this cookie
			Path:     "/",
		})
		http.Redirect(writer, request, "/dashboard", http.StatusSeeOther)
	} else {
		templ.Handler(view.LoginPage("Invalid email or password")).ServeHTTP(writer, request)
	}
}
