package server

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/middleware"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func CreateAPIServer(conn database.Database) (*echo.Echo, error) {
	e := echo.New()
	e.Server.Addr = configs.APIPort
	e.Logger.SetLevel(log.WARN)
	e.HTTPErrorHandler = handlers.HTTPErrorHandler

	e.Pre(middleware.AccessLogMiddleware)
	e.Pre(echoMid.AddTrailingSlash())
	e.Pre(middleware.ExpireInvalidCookiesMiddleware(conn))

	api := e.Group("/api")

	users := api.Group("/users")
	users.GET("/", handlers.GetHandlerUsers(conn))
	users.GET("/:user_id/", handlers.GetHandlerUser(conn))
	users.POST("/", handlers.GetHandlerUsersCreate(conn))
	users.PUT("/", handlers.GetHandlerUsersUpdate(conn))
	users.POST("/images/", handlers.GetImagesHandler(conn))

	films := api.Group("/films")
	films.GET("/:film_id/", handlers.GetHandlerFilm(conn))
	films.POST("/", handlers.GetHandlerFilmCreate(conn))
	films.POST("/images/", handlers.GetImagesHandler(conn))
	//films.DELETE("/:film_id", handlers.GetHandlerFilmDelete(conn))

	sessions := api.Group("/sessions")
	sessions.GET("/", handlers.GetHandlerLoginCheck(conn))
	sessions.POST("/", handlers.GetHandlerLogin(conn))
	sessions.DELETE("/", handlers.GetHandlerLogout(conn))

	e.Use(middleware.HeaderMiddleware)
	e.Use(echoMid.Recover())

	return e, nil
}
