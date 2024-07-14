package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				app.log.Error("Couldn't parse token from cookie.")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			app.log.Error("Something else happened parsing the cookie.")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tokenString := c.Value

		tkn, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Check to see if alg is same

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				app.log.Error(err.Error())
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			app.log.Error(err.Error())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			app.log.Error("Invalid token.")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Here we're going to send the userid obtained from the token
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", "test")
		r = r.WithContext(ctx)

		app.log.Info(fmt.Sprintf("User name is: %s", r.Context().Value("username")))
		next.ServeHTTP(w, r)
	})
}
