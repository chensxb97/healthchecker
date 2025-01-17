package web

import (
	"encoding/json"
	"healthchecker/internal/health"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	HealthChecker *health.HealthChecker
}

func NewServer(hc *health.HealthChecker) *Server {
	return &Server{HealthChecker: hc}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/status", s.statusHandler)
	mux.HandleFunc("/endpoint1", s.randomResponse)
	mux.HandleFunc("/endpoint2", s.randomResponse)
	mux.HandleFunc("/endpoint3", s.randomResponse)

	// Add CORS middleware
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow only the React app
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

	http.ListenAndServe(":8080", handler)
}

func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	statuses := s.HealthChecker.GetStatuses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (s *Server) randomResponse(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 { // Random number between 0 and 1
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
}
