package server

import (
	"log"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/middleware"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/interfaces"
	handlers "github.com/go-park-mail-ru/2019_2_Pirogi/app/interfaces/http"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

func CreateLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	cfg.OutputPaths = []string{
		"stdout",
		configs.Default.AccessLog,
	}
	cfg.ErrorOutputPaths = []string{
		"stderr",
		configs.Default.ErrorLog,
	}
	return cfg.Build()
}

func CreateAPIServer(conn database.Database) (*echo.Echo, error) {
	validation.InitValidator()
	e := echo.New()
	logger, err := CreateLogger()
	zap.ReplaceGlobals(logger)
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	e.Server.Addr = configs.Default.APIPort
	e.HTTPErrorHandler = handlers.GetHTTPErrorHandler(logger)

	e.Pre(middleware.GetAccessLogMiddleware(logger))
	e.Pre(echoMid.AddTrailingSlash())
	//e.Pre(middleware.PostCheckMiddleware)

	filmRepo := interfaces.NewFilmRepository(conn)
	personRepo := interfaces.NewPersonRepository(conn)
	userRepo := interfaces.NewUserRepository(conn)
	cookieRepo := interfaces.NewCookieRepository(conn)
	reviewRepo := interfaces.NewReviewRepository(conn)
	subscriptionRepo := interfaces.NewSubscriptionRepository(conn)

	filmUsecase := usecase.NewFilmUsecase(filmRepo, personRepo, subscriptionRepo)
	searchUsecase := usecase.NewSearchUsecase(filmRepo, personRepo)
	authUsecase := usecase.NewAuthUsecase(userRepo, cookieRepo, subscriptionRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, cookieRepo)
	personUsecase := usecase.NewPersonUsecase(personRepo, filmRepo, cookieRepo, subscriptionRepo)
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo, cookieRepo, userRepo)
	pagesUsecase := usecase.NewPagesUsecase(filmRepo, personRepo)
	imageUsecase := usecase.NewImageUsecase(cookieRepo, userRepo)
	subscriptionUsecase := usecase.NewSubscriptionUsecase(subscriptionRepo, cookieRepo, personRepo, userRepo)

	api := e.Group("/api")

	films := api.Group("/films")
	films.GET("/:film_id/", handlers.GetHandlerFilm(filmUsecase))
	films.POST("/", handlers.GetHandlerFilmCreate(filmUsecase))

	api.GET("/search/", handlers.GetHandlerSearch(searchUsecase))

	sessions := api.Group("/sessions")
	sessions.GET("/", handlers.GetHandlerLoginCheck(authUsecase))
	sessions.POST("/", handlers.GetHandlerLogin(authUsecase))
	sessions.DELETE("/", handlers.GetHandlerLogout(authUsecase))

	users := api.Group("/users")
	users.GET("/", handlers.GetHandlerUsers(userUsecase))
	users.GET("/:user_id/", handlers.GetHandlerUser(userUsecase))
	users.POST("/", handlers.GetHandlerUsersCreate(userUsecase))
	users.PUT("/", handlers.GetHandlerUsersUpdate(userUsecase))

	users.POST("/images/", handlers.GetImagesHandler(imageUsecase))

	persons := api.Group("/persons")
	persons.GET("/:person_id/", handlers.GetHandlerPerson(personUsecase))

	reviews := api.Group("/reviews")
	reviews.GET("/:film_id/", handlers.GetHandlerReviews(reviewUsecase))
	reviews.GET("/", handlers.GetHandlerProfileReviews(reviewUsecase))
	reviews.POST("/", handlers.GetHandlerReviewsCreate(reviewUsecase))

	subscriptions := api.Group("/subscriptions")
	subscriptions.GET("/", handlers.GetHandlerSubscriptionList(subscriptionUsecase))
	subscriptions.GET("/events/", handlers.GetHandlerNewEvents(subscriptionUsecase))
	subscriptions.DELETE("/events/", handlers.GetHandlerReadNewEvents(subscriptionUsecase))
	subscriptions.POST("/", handlers.GetHandlerSubscribe(subscriptionUsecase))
	subscriptions.DELETE("/", handlers.GetHandlerUnsubscribe(subscriptionUsecase))

	//likes := api.Group("/likes")
	//likes.POST("/", handlers.GetHandlerLikesCreate(conn))
	//
	//marks := api.Group("/marks")
	//marks.POST("/", handlers.GetHandlerRatingsCreate(conn))
	//
	//lists := api.Group("/lists")
	//lists.GET("/", handlers.GetHandlerList(conn))

	api.GET("/common/:variable/", handlers.HandlerCommon())

	pages := api.Group("/pages")
	pages.GET("/", handlers.GetHandlerPages(pagesUsecase))

	e.Use(echoMid.CORSWithConfig(echoMid.CORSConfig{
		AllowOrigins:     []string{"https://cinsear.ru", "http://localhost:8080"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	e.Use(echoMid.Secure())
	e.Use(middleware.SetCSRFCookie)
	e.Use(middleware.HeaderMiddleware)
	e.Use(echoMid.Recover())

	return e, nil
}
