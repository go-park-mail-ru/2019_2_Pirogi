package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("p", "8080", "port for server")
	flag.Parse()

	db := inmemory.Init()
	db.FakeFillDB()

	apiRouter := mux.NewRouter()

	apiRouter.Use(middleware.HeaderMiddleware)
	apiRouter.Use(middleware.LoggingMiddleware)
	apiRouter.Use(middleware.GetCheckAuthMiddleware(db))
	apiRouter.HandleFunc("/api/films/{film_id:[0-9]+}", handlers.GetHandlerFilm(db)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/api/users/", handlers.GetHandlerUsersCreate(db)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/api/users/", handlers.GetHandlerUsers(db)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/api/users/", handlers.GetHandlerUsersUpdate(db)).Methods(http.MethodPut)
	apiRouter.HandleFunc("/api/users/{user_id:[0-9]+}", handlers.GetHandlerUser(db)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/api/login/", handlers.GetHandlerLoginCheck(db)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/api/login/", handlers.GetHandlerLogin(db)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/api/logout/", handlers.GetHandlerLogout(db)).Methods(http.MethodPost)

	s := server.New(*port)
	s.Init(db, apiRouter)

	log.Printf("Starting api server at port :%s...", *port)
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
