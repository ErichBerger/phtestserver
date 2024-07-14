package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, pageName string, data data) {

	// Access the page from the cache
	templateSet, ok := app.templateCache[pageName]
	// If it doesn't exist, warn
	if !ok {

		app.log.Warn("Template not found", "name", pageName)
		return
	}

	w.WriteHeader(status)

	err := templateSet.ExecuteTemplate(w, "base", data)

	if err != nil {
		app.log.Warn("template not executed", "name", pageName)
		return
	}

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
