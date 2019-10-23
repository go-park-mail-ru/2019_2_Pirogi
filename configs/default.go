package configs

import "time"

const (
	CookieAuthName     = "cinsear_session"
	CookieAuthDuration = 10 * time.Hour

	AccessLog = "/log/access.log"
	ErrorLog  = "/log/error.log"

	UsersImageUploadPath = "/media/images/users/"
	FilmsImageUploadPath = "/media/images/films/"

	// Put the credentials of MongoDB user here
	MongoUser              = ""
	MongoPwd               = ""
	MongoHost              = "mongodb://127.0.0.1:27017"
	MongoDbName            = "cinsear"
	UsersCollectionName    = "users"
	FilmsCollectionName    = "films"
	CookiesCollectionName  = "cookies"
	CountersCollectionName = "counters"

	UserTargetName   = "user"
	FilmTargetName   = "film"
	CookieTargetName = "cookie"

	APIPort = ":8000"

	CertFile = "/ssl/cert.pem"
	KeyFile  = "/ssl/privkey.pem"
)
