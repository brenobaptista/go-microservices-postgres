package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/brenobaptista/go-microservices-postgres/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

func NewAPI(pgdb *pg.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	router.Route("/homes", func(subrouter chi.Router) {
		subrouter.Post("/", createHome)
		subrouter.Get("/{homeID}", getHomeByID)
		subrouter.Get("/", getHomes)
		subrouter.Put("/{homeID}", updateHomeByID)
		subrouter.Delete("/{homeID}", deleteHomeByID)
	})

	return router
}

type HomeRequest struct {
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
}

type HomeResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Home    *db.Home `json:"home"`
}

func createHome(w http.ResponseWriter, r *http.Request) {
	req := &HomeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get database from context",
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	home, err := db.CreateHome(pgdb, &db.Home{
		Price:       req.Price,
		Description: req.Description,
		Address:     req.Address,
		AgentID:     req.AgentID,
	})
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func getHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get database from context",
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	home, err := db.GetHome(pgdb, homeID)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

type HomesResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error"`
	Homes   []*db.Home `json:"homes"`
}

func getHomes(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get database from context",
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	homes, err := db.GetHomes(pgdb)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &HomesResponse{
		Success: true,
		Error:   "",
		Homes:   homes,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func updateHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	req := &HomeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get database from context",
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intHomeID, err := strconv.ParseInt(homeID, 10, 64)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	home, err := db.UpdateHome(pgdb, &db.Home{
		ID:          intHomeID,
		Price:       req.Price,
		Description: req.Description,
		Address:     req.Address,
		AgentID:     req.AgentID,
	})
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func deleteHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get database from context",
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intHomeID, err := strconv.ParseInt(homeID, 10, 64)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.DeleteHome(pgdb, intHomeID)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    nil,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}
