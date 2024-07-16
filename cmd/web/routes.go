package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//This is the same as mux.Handle("/", http.Handler(app.mainHandler))
	mux.HandleFunc("/", app.mainHandler)
	mux.HandleFunc("/progressnote", app.progressnoteHandler)
	mux.HandleFunc("/notes-admin", app.notesAdmin)
	mux.HandleFunc("/note-admin/{id}", app.noteAdmin)
	mux.Handle("/note/{id}", app.providerVerify(http.HandlerFunc(app.noteHandler)))

	//protected
	mux.Handle("/notes", app.providerVerify(http.HandlerFunc(app.notesHandler)))

	//Progress Notes Path
	mux.HandleFunc("/add-note-1", app.addNote1)
	mux.HandleFunc("/add-note-2", app.addNote2)
	mux.HandleFunc("/add-note-3", app.addNote3)
	mux.HandleFunc("/add-note-4", app.addNote4)
	mux.HandleFunc("/add-note-5", app.addNote5)
	mux.HandleFunc("/add-note-6", app.addNote6)
	mux.HandleFunc("/add-note-7", app.addNote7)
	mux.Handle("POST /add-note", app.providerVerify(http.HandlerFunc(app.addNotePost)))
	mux.Handle("GET /add-note", app.providerVerify(http.HandlerFunc(app.addNoteGet)))
	mux.Handle("/logout", app.providerVerify(http.HandlerFunc(app.logout)))
	//Login
	mux.HandleFunc("POST /auth", app.loginHandler)

	// We're using a closure over commonHeaders.
	// There are certain things we want to respond with not matter the request, so this
	// provides a way to call the commonHeaders function, which in turn calls the appropriate
	// function from mux

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
