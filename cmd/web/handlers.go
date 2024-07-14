package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		app.log.Error(err.Error())
		return
	}

	note, err := app.notes.Get(id)

	if err != nil {
		app.log.Error(err.Error())
		http.Error(w, "Note not found.", http.StatusBadRequest)
		return
	}

	data := data{
		Note: note,
	}

	app.render(w, r, http.StatusOK, "note.html", data)
}

func (app *application) addNotePost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	patient := r.PostForm.Get("patient")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse patient")
	}
	service := r.PostForm.Get("service")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse service")
	}
	serviceDate := r.PostForm.Get("serviceDate")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse serviceDate")
	}
	startTime := r.PostForm.Get("startTime")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse startTime")
	}
	endTime := r.PostForm.Get("endTime")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse endTime")
	}
	summary := r.PostForm.Get("summary")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse summary")
	}
	progress := r.PostForm.Get("progress")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse progress")
	}
	response := r.PostForm.Get("response")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse response")
	}
	assessmentStatus := r.PostForm.Get("assessmentStatus")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse assessmentStatus")
	}
	riskFactors := r.PostForm.Get("riskFactors")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse riskFactors")
	}
	emergencyInterventions := r.PostForm.Get("emergencyInterventions")
	if strings.Compare("", patient) == 0 {
		app.log.Error("Couldn't parse emergencyInterventions")
	}

	// Get userID from context
	provider := r.Context().Value("username").(string)

	if strings.Compare(provider, "") == 0 {
		app.log.Error("Failed to parse providerID from context.")
		http.Error(w, "Couldn't get user ID from context", http.StatusInternalServerError)
		return
	}

	providerID, err := app.users.GetID(provider)

	if err != nil {

		app.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newID, err := app.notes.Insert(providerID, patient, service, serviceDate, startTime, endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions)

	if err != nil {
		app.log.Error("note insertion gone wrong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/%d", newID), http.StatusSeeOther)

	// upon successful entry, redirect to new note view

}

func (app *application) addNoteGet(w http.ResponseWriter, r *http.Request) {
	data := data{}

	username := r.Context().Value("username")

	if username == nil {
		app.log.Error("No context with username provided.")
	}

	app.log.Info(fmt.Sprintf("username: %s", username))

	app.render(w, r, http.StatusOK, "add-note.html", data)
}

func (app *application) addNote1(w http.ResponseWriter, r *http.Request) {

	data := data{}

	app.render(w, r, http.StatusOK, "add-note-1.html", data)
}

func (app *application) addNote2(w http.ResponseWriter, r *http.Request) {

	// If no note in session, redirect to first note page

	// If note in session with data relevant to page, populate form with data

	// If note in session with no relevant data, do not populate form
	app.render(w, r, http.StatusOK, "add-note-2.html", data{})
}

func (app *application) addNote2Post(w http.ResponseWriter, r *http.Request) {
	// Same as addNote2, but handles if data has been submitted from addNote1
	// Adds to session data for ongoing note
	// If data fields already exist for 2, does the same thing as addNote2
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
