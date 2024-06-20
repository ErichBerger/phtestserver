package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	templateCache map[string]*template.Template
	log           *slog.Logger
}

func main() {

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tc, err := newTemplateCache()

	if err != nil {
		log.Warn(err.Error())
		os.Exit(1)
	}

	app := &application{
		templateCache: tc,
		log:           log,
	}

	srv := &http.Server{
		Addr:        ":8080",
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
		ErrorLog:    slog.NewLogLogger(log.Handler(), slog.LevelInfo),
	}

	log.Info("starting server", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		app.log.Warn(err.Error())
		os.Exit(1)
	}
}
