package main

import (
	"database/sql"
	"errors"
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
		app.serverError(w, r, err)
		return
	}

	note, err := app.notes.Get(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		app.serverError(w, r, err)
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
	if strings.TrimSpace(patient) == "" {
		app.serverError(w, r, fmt.Errorf("patient field not submitted"))
	}

	service := r.PostForm.Get("service")
	if strings.TrimSpace(service) == "" {
		app.serverError(w, r, fmt.Errorf("service field not submitted"))
	}

	serviceDate := r.PostForm.Get("serviceDate")
	if strings.TrimSpace(serviceDate) == "" {
		app.serverError(w, r, fmt.Errorf("serviceDate field not submitted"))
	}

	startTime := r.PostForm.Get("startTime")
	if strings.TrimSpace(startTime) == "" {
		app.serverError(w, r, fmt.Errorf("startTime field not submitted"))
	}

	endTime := r.PostForm.Get("endTime")
	if strings.TrimSpace(endTime) == "" {
		app.serverError(w, r, fmt.Errorf("endTime field not submitted"))
	}

	summary := r.PostForm.Get("summary")
	if strings.TrimSpace(summary) == "" {
		app.serverError(w, r, fmt.Errorf("summary field not submitted"))
	}

	progress := r.PostForm.Get("progress")
	if strings.TrimSpace(progress) == "" {
		app.serverError(w, r, fmt.Errorf("progress field not submitted"))
	}

	response := r.PostForm.Get("response")
	if strings.TrimSpace(response) == "" {
		app.serverError(w, r, fmt.Errorf("response field not submitted"))
	}

	assessmentStatus := r.PostForm.Get("assessmentStatus")
	if strings.TrimSpace(assessmentStatus) == "" {
		app.serverError(w, r, fmt.Errorf("assessmentStatus field not submitted"))
	}

	riskFactors := r.PostForm.Get("riskFactors")
	if strings.TrimSpace(riskFactors) == "" {
		app.serverError(w, r, fmt.Errorf("riskFactors field not submitted"))
	}

	emergencyInterventions := r.PostForm.Get("emergencyInterventions")
	if strings.TrimSpace(emergencyInterventions) == "" {
		app.serverError(w, r, fmt.Errorf("emergencyInterventions field not submitted"))
	}

	// Get userID from context
	provider := r.Context().Value("username").(string)

	if strings.Compare(provider, "") == 0 {
		app.serverError(w, r, fmt.Errorf("failed to parse provider from context"))
		return
	}

	providerID, err := app.users.GetID(provider)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	newID, err := app.notes.Insert(providerID, patient, service, serviceDate, startTime, endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions)

	if err != nil {
		app.serverError(w, r, err)
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
