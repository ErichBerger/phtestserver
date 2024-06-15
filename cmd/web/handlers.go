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

func (app *application) progressnoteHandler1(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-1.html")
}

func (app *application) progressnoteHandler2(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-2.html")
}
func (app *application) progressnoteHandler3(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-3.html")
}
func (app *application) progressnoteHandler4(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-4.html")
}
func (app *application) progressnoteHandler5(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-5.html")
}
func (app *application) progressnoteHandler6(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-6.html")
}
func (app *application) progressnoteHandler7(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote-7.html")
}
