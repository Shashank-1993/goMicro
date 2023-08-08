package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app Config)routes() http.Handler {

	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedHeaders: []string{"Accept","Authorization","Content-Type","X-CSRF-Token"},
		AllowedMethods: []string{"POST","GET","PUT","DELETE","OPTIONS"},
		AllowCredentials: true,
		ExposedHeaders: []string{"Link"},
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/",app.Broker)

	return mux
}
