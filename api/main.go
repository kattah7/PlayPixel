package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kattah7/v3/models"
	"github.com/kattah7/v3/storage"
)

type apiFunc func(http.ResponseWriter, *http.Request, *APIServer) error

type APIServer struct {
	listenAddr string
	store      storage.Storage
	cfg        *models.Config
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewAPIServer(cfg *models.Config, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: cfg.ListenAddress,
		store:      store,
		cfg:        cfg,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = s.customHandler(route.HandlerFunc)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, http.StatusOK, ApiResponse{
			Success: false,
			Error:   "Endpoint not found",
		})
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, http.StatusOK, ApiResponse{
			Success: false,
			Error:   "Method not allowed",
		})
	})

	log.Println("JSON API server running on port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) customHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			s.WriteJSON(w, http.StatusOK, ApiResponse{
				Success: false,
				Error:   "No token provided",
			})
			return
		}

		if tokenString != s.cfg.Auth {
			s.WriteJSON(w, http.StatusOK, ApiResponse{
				Success: false,
				Error:   "Invalid token",
			})
			return
		}

		log.Println(r.Method, r.URL.Path)

		if err := f(w, r, s); err != nil {
			s.WriteJSON(w, http.StatusOK, ApiResponse{Error: err.Error()})
		}
	}
}

func (s *APIServer) WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
