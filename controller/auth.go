package controller

import (
	"CeratOps/model"
	"CeratOps/view"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// If GET show the page
	if r.Method == http.MethodGet {
		if model.IsAuthenticated(r) {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		templ.Handler(view.LoginPage("")).ServeHTTP(w, r)
		return
	}

	// If POST process the login
	email := r.FormValue("email")
	pass := r.FormValue("password")
	user := model.GetUser()

	// Check email and use internal model utility for password verification
	if email == user.Email && model.CheckPasswordHash(pass, user.PasswordHash) {
		token, err := model.GenerateJWT(email)
		if err != nil {
			templ.Handler(view.LoginPage("Error generating token")).ServeHTTP(w, r)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true, // JS can't read this cookie
			Path:     "/",
		})
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		templ.Handler(view.LoginPage("Invalid email or password")).ServeHTTP(w, r)
	}
}
