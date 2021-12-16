package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewAPI() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/homes", func(r chi.Router) {
		r.Post("/", createHome)
		r.Get("/{homeID}", getHomeByID)
		r.Get("/", getHomes)
		r.Put("/{homeID}", updateHomeByID)
		r.Delete("/{homeID}", deleteHomeByID)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!\n"))
	})

	return r
}

func createHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create home\n"))
}

func getHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	w.Write([]byte(fmt.Sprintf("get home: %s\n", homeID)))
}

func getHomes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all homes\n"))
}

func updateHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	w.Write([]byte(fmt.Sprintf("update home: %s\n", homeID)))
}

func deleteHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	w.Write([]byte(fmt.Sprintf("delete home: %s\n", homeID)))
}
