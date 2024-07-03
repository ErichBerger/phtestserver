package main

import (
	"net/http"
)

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set some common headers here
		w.Header().Set("server", "Go")

		next.ServeHTTP(w, r)

		// anything here will be called only after the next handler's method has finished
	})
}

func (app *application) adminCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement some logic to check if

		valid, err := app.users.ValidateAdmin()

		if err != nil {
			// Set internal server error
			return
		}

		if !valid {
			// Set response as some access restricted error
			return
		}
		next.ServeHTTP(w, r)
	})
}
