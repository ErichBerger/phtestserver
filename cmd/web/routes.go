package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//This is the same as mux.Handle("/", http.Handler(app.mainHandler))
	mux.HandleFunc("/", app.mainHandler)
	mux.HandleFunc("/colortest", app.colorTestHandler)
	mux.HandleFunc("/progressnote", app.progressnoteHandler)
	mux.HandleFunc("/admindashboard", app.admindashboardHandler)
	mux.HandleFunc("/adminnoteview", app.adminnoteviewHandler)
	mux.HandleFunc("/providerdashboard", app.providerdashboardHandler)
	return mux
}
