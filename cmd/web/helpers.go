package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.log.Error(err.Error(), "method", method, "uri", uri)

	app.clientError(w, http.StatusInternalServerError)

}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, pageName string, data data) {

	// Access the page from the cache
	templateSet, ok := app.templateCache[pageName]
	// If it doesn't exist, warn
	if !ok {
		err := fmt.Errorf("the template %s does not exist", pageName)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := templateSet.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	buf.WriteTo(w)

}

func (app *application) HTMLTimeToGoTime(inputDate string, inputTime string) (time.Time, error) {

	dates := strings.Split(inputDate, "-")

	year, err := strconv.Atoi(dates[0])

	if err != nil {
		return time.Time{}, err
	}

	month, err := strconv.Atoi(dates[1])

	if err != nil {

		return time.Time{}, err
	}

	day, err := strconv.Atoi(dates[2])

	if err != nil {

		return time.Time{}, err
	}

	times := strings.Split(inputTime, ":")

	hour, err := strconv.Atoi(times[0])

	if err != nil {

		return time.Time{}, err
	}

	minute, err := strconv.Atoi(times[1])

	if err != nil {

		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC), nil
}

func (app *application) getTemplateData(r *http.Request) data {
	username := r.Context().Value("username").(string)
	if username == "" {
		return data{IsLoggedIn: false}
	}

	return data{IsLoggedIn: true}
}
