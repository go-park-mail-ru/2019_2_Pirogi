package server

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/validators"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"log"
)

func CreateLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"stdout",
		"/log/cinsear.log",
	}
	cfg.ErrorOutputPaths = []string{
		"stderr",
		"/log/error.log",
	}
	return cfg.Build()
}

func CreateAPIServer(conn database.Database) (*echo.Echo, error) {
	validators.InitValidator()

	e := echo.New()
	logger, err := CreateLogger()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	e.Server.Addr = configs.Default.APIPort
	e.HTTPErrorHandler = handlers.GetHTTPErrorHandler(logger)

	e.Pre(middleware.GetAccessLogMiddleware(logger))
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

	persons := api.Group("/persons")
	persons.GET("/:person_id/", handlers.GetHandlerPerson(conn))
	persons.POST("/", handlers.GetHandlerPersonsCreate(conn))
	persons.POST("/images/", handlers.GetImagesHandler(conn))

	reviews := api.Group("/reviews")
	reviews.GET("/:film_id/", handlers.GetHandlerReviews(conn))
	reviews.GET("/", handlers.GetHandlerProfileReviews(conn))
	reviews.POST("/", handlers.GetHandlerReviewsCreate(conn))

	likes := api.Group("/likes")
	likes.POST("/", handlers.GetHandlerLikesCreate(conn))

	marks := api.Group("/marks")
	marks.POST("/", handlers.GetHandlerRatingsCreate(conn))

	lists := api.Group("/lists")
	lists.GET("/", handlers.GetHandlerList(conn))

	common := api.Group("/common")
	common.GET("/:variable/", handlers.HandlerCommon())

	pages := api.Group("/pages")
	pages.GET("/", handlers.GetHandlerPages(conn))

	e.Use(echoMid.CORSWithConfig(echoMid.CORSConfig{
		AllowOrigins: []string{"https://cinsear.ru", "http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(echoMid.Secure())
	e.Use(middleware.SetCSRFCookie)
	e.Use(middleware.HeaderMiddleware)
	e.Use(echoMid.Recover())

	return e, nil
}
