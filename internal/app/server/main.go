package server

import (
	"../../pkg/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const base = "/api/"
const timeout = 10

type Server struct {
	port    string
	handler mux.Router
}

func New(port string) Server {
	s := Server{port: port}
	return s
}

func (s *Server) Init() {
	apiRouter := mux.NewRouter()

	apiRouter.HandleFunc("/api/users/{user_id}", handlers.HandleUser).Methods(http.MethodGet)
	apiRouter.HandleFunc("/api/users/", handlers.HandleUsers).Methods(http.MethodPost, http.MethodPut, http.MethodDelete)
	apiRouter.HandleFunc("/", handlers.HandleDefault)

	s.handler = *apiRouter
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         ":" + s.port,
		Handler:      &s.handler,
		ReadTimeout:  timeout * time.Second,
		WriteTimeout: timeout * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
