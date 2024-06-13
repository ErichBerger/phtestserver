package main

import (
	"net/http"
)

func (app *application) mainHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html")
}

func (app *application) progressnoteHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote.html")
}

func (app *application) colorTestHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "colortest.html")
}

func (app *application) admindashboardHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "admindashboard.html")
}
func (app *application) adminnoteviewHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "adminnoteview.html")
}
func (app *application) providerdashboardHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "providerdashboard.html")
}
