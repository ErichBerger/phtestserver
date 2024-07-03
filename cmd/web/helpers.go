package main

import (
	"net/http"
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
