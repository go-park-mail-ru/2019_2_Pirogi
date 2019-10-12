package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/gorilla/mux"
)

func CreateAPIServer(port string, db database.Database) server.Server {
	router := mux.NewRouter()
	router.Use(middleware.HeaderMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.GetCheckAuthMiddleware(db))

	subrouter := router.PathPrefix("/api").Subrouter()

	subrouter.HandleFunc("/users/images/", handlers.GetUploadImageHandler(db, "users")).Methods(http.MethodPost)
	subrouter.HandleFunc("/films/images/", handlers.GetUploadImageHandler(db, "films")).Methods(http.MethodPost)

	subrouter.HandleFunc("/films/{film_id:[0-9]+}/", handlers.GetHandlerFilm(db)).Methods(http.MethodGet)

	subrouter.HandleFunc("/users/", handlers.GetHandlerUsersCreate(db)).Methods(http.MethodPost)
	subrouter.HandleFunc("/users/", handlers.GetHandlerUsers(db)).Methods(http.MethodGet)
	subrouter.HandleFunc("/users/{user_id:[0-9]+}/", handlers.GetHandlerUser(db)).Methods(http.MethodGet)
	subrouter.HandleFunc("/users/", handlers.GetHandlerUsersUpdate(db)).Methods(http.MethodPut)

	subrouter.HandleFunc("/sessions/", handlers.GetHandlerLoginCheck(db)).Methods(http.MethodGet)
	subrouter.HandleFunc("/sessions/", handlers.GetHandlerLogin(db)).Methods(http.MethodPost)
	subrouter.HandleFunc("/sessions/", handlers.GetHandlerLogout(db)).Methods(http.MethodDelete)

	s := server.New(port)
	s.Init(router)
	return s
}

func main() {
	portAPI := flag.String("api", "8000", "port for API server")
	flag.Parse()

	conn, err := database.InitMongo()
	if err != nil {
		log.Fatal(err)
	}

	// Do it one time
	conn.FakeFillDB()

	apiServer := CreateAPIServer(*portAPI, conn)
	apiServer.Run()
}
