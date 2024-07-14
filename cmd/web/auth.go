package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtKey = []byte("secret-key")
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.log.Error("Couldn't parse form.")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	username := r.PostForm.Get("user")
	if strings.Compare(username, "") == 0 {
		app.log.Error("Failed to parse username.")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	password := r.PostForm.Get("pass")
	if strings.Compare(password, "") == 0 {
		app.log.Error("Failed to parse password.")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	// Get hashed password from DB and compare with submitted hashed password
	storedPasswordHash := "$2y$10$X8XV2SPQ4sVyYqCXpmTTlucH3QLqm7lStxkY4jjQQxuj5yV8WfMzm" // bcrypt hash for "password"

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))

	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime), Subject: username})
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}
