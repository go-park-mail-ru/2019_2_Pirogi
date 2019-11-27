package server

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/metrics"
	"log"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	v12 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/sessions/protobuf"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/middleware"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/interfaces"
	handlers "github.com/go-park-mail-ru/2019_2_Pirogi/app/interfaces/http"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func CreateAPIServer(conn database.Database) (*echo.Echo, error) {
	validation.InitValidator()
	metrics.InitMetrics()
	e := echo.New()
	logger, err := network.CreateLogger()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	usersConn, err := grpc.Dial(
		"users"+configs.Default.UsersMicroservicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		zap.S().Error(err.Error())
	} else {
		zap.S().Debug("Dialing with users server was successful")
	}

	sessionsConn, err := grpc.Dial(
		"sessions"+configs.Default.SessionsMicroservicePort,
		grpc.WithInsecure())
	if err != nil {
		zap.S().Error(err.Error())
	} else {
		zap.S().Debug("Dialing with sessions server was successful")
	}

	sessionsClient := v12.NewAuthServiceClient(sessionsConn)
	usersClient := v1.NewUserServiceClient(usersConn)

	e.Server.Addr = configs.Default.APIPort
	e.HTTPErrorHandler = handlers.GetHTTPErrorHandler(logger)

	e.Pre(middleware.GetAccessLogMiddleware(logger))
	e.Pre(echoMid.AddTrailingSlash())

	filmRepo := interfaces.NewFilmRepository(conn)
	personRepo := interfaces.NewPersonRepository(conn)
	userRepo := interfaces.NewUserRepository(conn)
	cookieRepo := interfaces.NewCookieRepository(conn)
	reviewRepo := interfaces.NewReviewRepository(conn)
	subscriptionRepo := interfaces.NewSubscriptionRepository(conn)

	filmUsecase := usecase.NewFilmUsecase(filmRepo, personRepo, subscriptionRepo)
	searchUsecase := usecase.NewSearchUsecase(filmRepo, personRepo)
	authUsecase := usecase.NewAuthUsecase(subscriptionRepo, sessionsClient)
	userUsecase := usecase.NewUserUsecase(usersClient, sessionsClient)
	personUsecase := usecase.NewPersonUsecase(personRepo, filmRepo, cookieRepo, subscriptionRepo)
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo, cookieRepo, userRepo)
	pagesUsecase := usecase.NewPagesUsecase(filmRepo, personRepo)
	imageUsecase := usecase.NewImageUsecase(cookieRepo, userRepo)
	subscriptionUsecase := usecase.NewSubscriptionUsecase(subscriptionRepo, cookieRepo, personRepo, userRepo)

	e.GET("/metrics/", echo.WrapHandler(promhttp.Handler()))

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
	e.Use(middleware.GetMetricsMiddleware)
	e.Use(echoMid.Recover())

	return e, nil
}
