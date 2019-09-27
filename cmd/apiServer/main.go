package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port for server")
	flag.Parse()

	db := inmemory.Init()
	db.FakeFillDB()

	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/api/users/", handlers.GetHandlerUsersCreate(db)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/api/users/{user_id:[0-9]+}", handlers.GetHandlerUser(db)).Methods(http.MethodGet)
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
