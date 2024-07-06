package main

import (
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtKey = []byte("secret-key")
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	username := r.FormValue("user")
	password := r.FormValue("pass")

	storedPasswordHash := "$2y$10$X8XV2SPQ4sVyYqCXpmTTlucH3QLqm7lStxkY4jjQQxuj5yV8WfMzm" // bcrypt hash for "password"

	err := bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))

	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}