package api

import (
	"log"
	"net/http"
)

func (s *APIServer) Middleware(f apiFunc, auth bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				s.WriteJSON(w, http.StatusOK, ApiResponse{
					Success: false,
					Error:   "No token provided",
				})
				return
			}

			if tokenString != s.cfg.Authorization {
				s.WriteJSON(w, http.StatusOK, ApiResponse{
					Success: false,
					Error:   "Invalid token",
				})
				return
			}
		}

		log.Println(r.Method, r.URL.Path)

		if err := f(w, r); err != nil {
			s.WriteJSON(w, http.StatusOK, ApiResponse{Error: err.Error()})
		}
	}
}
