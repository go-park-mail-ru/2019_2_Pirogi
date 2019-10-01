package main

import (
	"flag"
	"net/http"
	"sync"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/gorilla/mux"
)

func CreateAPIServer(port string, db *inmemory.DB) server.Server {
	router := mux.NewRouter()
	router.Use(middleware.HeaderMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.GetCheckAuthMiddleware(db))

	router.HandleFunc("/api/films/{film_id:[0-9]+}", handlers.GetHandlerFilm(db)).Methods(http.MethodGet)

	router.HandleFunc("/api/users/", handlers.GetHandlerUsersCreate(db)).Methods(http.MethodPost)
	router.HandleFunc("/api/users/", handlers.GetHandlerUsers(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{user_id:[0-9]+}/", handlers.GetHandlerUser(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/users/", handlers.GetHandlerUsersUpdate(db)).Methods(http.MethodPut)

	router.HandleFunc("/api/sessions/", handlers.GetHandlerLoginCheck(db)).Methods(http.MethodGet)
	router.HandleFunc("/api/sessions/", handlers.GetHandlerLogin(db)).Methods(http.MethodPost)
	router.HandleFunc("/api/sessions/", handlers.GetHandlerLogout(db)).Methods(http.MethodDelete)

	router.HandleFunc("/api/users/images/", handlers.GetUploadImageHandler(db, "users")).Methods(http.MethodPost)
	router.HandleFunc("/api/films/images/", handlers.GetUploadImageHandler(db, "films")).Methods(http.MethodPost)

	s := server.New(port)
	s.Init(db, router)
	return s
}

func main() {
	portAPI := flag.String("api", "8000", "port for API server")
	flag.Parse()

	db := inmemory.Init()
	db.FakeFillDB()

	wg := &sync.WaitGroup{}
	apiServer := CreateAPIServer(*portAPI, db)
	wg.Add(2)
	go apiServer.Run(wg)
	wg.Wait()
}
