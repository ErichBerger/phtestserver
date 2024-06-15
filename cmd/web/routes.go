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

	//Progress Notes Path
	mux.HandleFunc("/progressnote-1", app.progressnoteHandler1)
	mux.HandleFunc("/progressnote-2", app.progressnoteHandler2)
	mux.HandleFunc("/progressnote-3", app.progressnoteHandler3)
	mux.HandleFunc("/progressnote-4", app.progressnoteHandler4)
	mux.HandleFunc("/progressnote-5", app.progressnoteHandler5)
	mux.HandleFunc("/progressnote-6", app.progressnoteHandler6)
	mux.HandleFunc("/progressnote-7", app.progressnoteHandler7)
	return mux
}
