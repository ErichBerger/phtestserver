package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data := app.getTemplateData(r)
		app.render(w, r, http.StatusOK, "home.html", data)

	})
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "login.html", data)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if errors.Is(err, http.ErrNoCookie) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	c.Expires = time.Now()

	// Override token cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) getAdminNotesView(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "notes-admin.html", data)
}

func (app *application) getAdminNoteView(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "note-admin.html", data)
}

func (app *application) getNotesView(w http.ResponseWriter, r *http.Request) {

	username, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("no username stored in context"))
		return
	}
	notes, err := app.notes.GetNotesByProvider(username)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.getTemplateData(r)

	data.Notes = notes

	app.render(w, r, http.StatusOK, "notes.html", data)
}

func (app *application) getNoteView(w http.ResponseWriter, r *http.Request) {
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
			app.clientError(w, r, http.StatusBadRequest)
			return
		}
		app.serverError(w, r, err)
		return
	}

	data := app.getTemplateData(r)

	data.Note = note
	/* TODO: reinstate code, but it's not recognizing the user with the note
	username, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("couldn't parse username from context"))
		return
	}


	providerID, err := app.users.GetID(username)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

		if data.Note.ProviderID != providerID {
			app.clientError(w, r, http.StatusForbidden)
			return
		}
	*/
	app.render(w, r, http.StatusOK, "note.html", data)
}

type NoteCreateForm struct {
	Patient     string
	Patients    []string
	Service     string
	Services    map[string]string
	ServiceDate string
	StartTime   string
	EndTime     string
	Summary     string
}

func (app *application) postNoteCreate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// Get patientID from initials
	// NOTE: doing it this way to prevent storage of patient id
	patient := r.PostForm.Get("patient")
	if strings.TrimSpace(patient) == "" {
		app.serverError(w, r, fmt.Errorf("patient field not submitted"))
	}

	patientInitials := strings.Split(patient, " ")

	patientFirst := patientInitials[0]

	patientLast := patientInitials[1]

	patientID, err := app.patients.GetID(patientFirst, patientLast)

	if err != nil {
		app.serverError(w, r, err)
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

	// Get username from context
	provider, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("failed to parse provider from context"))
		return
	}

	providerID, err := app.users.GetID(provider)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	newID, err := app.notes.Insert(providerID, patientID, service, serviceDate, startTime, endTime, summary)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/view/%d", newID), http.StatusSeeOther)

	// upon successful entry, redirect to new note view

}

func (app *application) getNoteCreate(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)

	/* Not sure why I had this orignally, maybe keep it around just in case
	username, ok := r.Context().Value("username").(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("error parsing username from context"))
	}

	userID, err := app.users.GetID(username)

	if err != nil {
		app.serverError(w, r, err)
	}

	*/
	form := NoteCreateForm{}

	// Get list of patients
	patients, err := app.patients.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	form.Patients = make([]string, 0)

	// Store their initials as a string and add to form data
	for _, patient := range patients {
		b := strings.Builder{}
		fmt.Fprintf(&b, "%s %s", patient.FirstInitials, patient.LastInitials)
		form.Patients = append(form.Patients, b.String())
	}
	form.Patient = "none"
	form.Services = map[string]string{
		"general":    "General",
		"individual": "Individual",
		"family":     "Family",
		"group":      "Group",
	}
	form.Service = "none"
	// Convert patients into form form
	data.Form = form

	app.render(w, r, http.StatusOK, "add-note.html", data)
}
