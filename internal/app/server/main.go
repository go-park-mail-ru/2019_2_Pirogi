package server

import (
	"../../pkg/handlers"
	"../../pkg/inmemory"
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
	db := inmemory.Init()
	db.FakeFillDB()

	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/api/users/", handlers.GetHandlerUsersCreate(db)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/api/users/{user_id:[0-9]+}", handlers.GetHandlerUser(db)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/api/login/", handlers.GetHandlerLogin(db)).Methods(http.MethodGet, http.MethodPost)
	apiRouter.HandleFunc("/api/logout/", handlers.GetHandlerLogout(db)).Methods(http.MethodPost)
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
