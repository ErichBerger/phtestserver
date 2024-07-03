package main

import (
	"net/http"
)

func (app *application) mainHandler(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, http.StatusOK, "home.html", data{})
}

func (app *application) progressnoteHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote.html", data{})
}

func (app *application) notesAdmin(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "notes-admin.html", data{})
}
func (app *application) noteAdmin(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "note-admin.html", data{})
}
func (app *application) notesHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "notes.html", data{})
}

func (app *application) noteHandler(w http.ResponseWriter, r *http.Request) {
	// We're assuming we've already checked if they have appropriate authorization
	// For now we're passing a dummy variable
	note, err := app.notes.Get(127)

	if err != nil {
		// log error, return not found
		return
	}

	data := data{
		Note: note,
	}

	app.render(w, r, http.StatusOK, "note.html", data)
}

func (app *application) addNote1(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-1.html", data{})
}

func (app *application) addNote2(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-2.html", data{})
}
func (app *application) addNote3(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-3.html", data{})
}
func (app *application) addNote4(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-4.html", data{})
}
func (app *application) addNote5(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-5.html", data{})
}
func (app *application) addNote6(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-6.html", data{})
}
func (app *application) addNote7(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-7.html", data{})
}
