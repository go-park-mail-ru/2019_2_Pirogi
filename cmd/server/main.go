package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func CreateAPIServer(conn database.Database) (*echo.Echo, error) {
	e := echo.New()

	e.Logger.SetLevel(log.WARN)
	e.HTTPErrorHandler = handlers.HTTPErrorHandler
	api := e.Group("/api")
	users := api.Group("/users")
	users.GET("/", handlers.GetHandlerUsers(conn))
	users.GET("/:user_id", handlers.GetHandlerUser(conn))
	users.POST("/", handlers.GetHandlerUsersCreate(conn))
	users.PUT("/", handlers.GetHandlerUsersUpdate(conn))
	//router.Use(middleware.HeaderMiddleware)
	//router.Use(middleware.LoggingMiddleware)
	//router.Use(middleware.GetCheckAuthMiddleware(conn))

	//subrouter.HandleFunc("/users/images/", handlers.GetUploadImageHandler(conn, "users")).Methods(http.MethodPost)
	//subrouter.HandleFunc("/films/images/", handlers.GetUploadImageHandler(conn, "films")).Methods(http.MethodPost)
	//
	//subrouter.HandleFunc("/films/{film_id:[0-9]+}/", handlers.GetHandlerFilm(conn)).Methods(http.MethodGet)
	//
	//subrouter.HandleFunc("/sessions/", handlers.GetHandlerLoginCheck(conn)).Methods(http.MethodGet)
	//subrouter.HandleFunc("/sessions/", handlers.GetHandlerLogin(conn)).Methods(http.MethodPost)
	//subrouter.HandleFunc("/sessions/", handlers.GetHandlerLogout(conn)).Methods(http.MethodDelete)

	return e, nil
}

func main() {
	portAPI := flag.String("api", "8000", "port for API server")
	flag.Parse()

	conn, err := database.InitInmemory()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	conn.FakeFillDB()

	apiServer, err := CreateAPIServer(conn)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Fatal(apiServer.Start("127.0.0.1:" + *portAPI))
}
