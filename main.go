package main

import (
	"CeratOps/controller"
	"CeratOps/model"
	"fmt"
	"net/http"
)

func main() {
	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Check auth and redirect
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if model.IsAuthenticated(r) {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})

	// Public Routes
	http.HandleFunc("/login", controller.LoginHandler)

	// Protected Routes (Wrapped in Middleware)
	http.HandleFunc("/dashboard", controller.AuthMiddleware(controller.DashboardHandler))

	fmt.Println("CeratOps running at http://localhost:8080")
	http.ListenAndServe(":8080", controller.PathMiddleware(http.DefaultServeMux))
}
