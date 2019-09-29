package server

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/gorilla/mux"
)

const timeout = 10

type Server struct {
	port    string
	handler mux.Router
}

func New(port string) Server {
	s := Server{port: port}
	return s
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

func (s *Server) Init(db *inmemory.DB, router *mux.Router) {
	db = inmemory.Init()
	apiRouter := router

	s.handler = *apiRouter
}
