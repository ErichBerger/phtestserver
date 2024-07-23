package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// Serving style and scripts as files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Public
	mux.HandleFunc("/", app.home)

	mux.Handle("/logout", app.providerVerify(http.HandlerFunc(app.logout)))

	mux.HandleFunc("POST /login", app.postLogin)

	mux.HandleFunc("GET /login", app.getLogin)

	// Admin
	mux.HandleFunc("/admin/notes/view", app.getAdminNotesView)

	mux.HandleFunc("/admin/note/view/{id}", app.getAdminNoteView)

	// Provider
	mux.Handle("/note/view/{id}", app.providerVerify(http.HandlerFunc(app.getNoteView)))

	mux.Handle("/notes/view", app.providerVerify(http.HandlerFunc(app.getNotesView)))

	mux.Handle("POST /note/create", app.providerVerify(http.HandlerFunc(app.postNoteCreate)))

	mux.Handle("GET /note/create", app.providerVerify(http.HandlerFunc(app.getNoteCreate)))

	// We're using a closure over commonHeaders.
	// There are certain things we want to respond with not matter the request, so this
	// provides a way to call the commonHeaders function, which in turn calls the appropriate
	// function from mux

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
