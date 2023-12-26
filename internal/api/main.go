package api

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/PlayPixel/api/internal/db"
	"github.com/PlayPixel/api/internal/logger"
	"github.com/PlayPixel/api/pkg/config"
	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	ctx   context.Context
	cfg   config.Config
	store db.Pool
	log   logger.Logger
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Init(ctx context.Context, cfg config.Config, store db.Pool, log logger.Logger) *APIServer {
	return &APIServer{
		ctx:   ctx,
		cfg:   cfg,
		store: store,
		log:   log,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range s.RoutesHandler() {
		handler := s.Middleware(route.HandlerFunc, route.RequireAuth)
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

	log.Println("PlayPixel API Listening on", s.cfg.BindAddress)

	srv := &http.Server{
		Handler:      router,
		Addr:         s.cfg.BindAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			return s.ctx
		},
	}

	log.Fatal(srv.ListenAndServe())
}

func (s *APIServer) WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
