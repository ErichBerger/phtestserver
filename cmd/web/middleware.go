package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set some common headers here
		w.Header().Set("Server", "Go")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		next.ServeHTTP(w, r)

		// anything here will be called only after the next handler's method has finished
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		proto := r.Proto
		method := r.Method
		uri := r.URL.RequestURI()

		app.log.Info("receivedrequest", "ip", ip, "proto", proto, "method", method, "uri", uri)
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) providerVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			app.serverError(w, r, err)
			return
		}

		tkn, err := jwt.ParseWithClaims(c.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Check to see if alg is same

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				app.clientError(w, http.StatusSeeOther)
				return
			}
			app.serverError(w, r, err)
			return
		}

		if !tkn.Valid {
			app.clientError(w, http.StatusUnauthorized)
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
