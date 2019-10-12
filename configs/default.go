package configs

import "time"

const CookieAuthName = "cinsear_session"
const CookieAuthDuration = 10 * time.Hour

const AccessLog = "/log/access.log"
const ErrorLog = "/log/error.log"

const UsersImageUploadPath = "/media/images/users/"
const FilmsImageUploadPath = "/media/images/films/"

const MongoDbUri = "mongodb://127.0.0.1:27017"
const MongoDbName = "cinsear"
const UsersCollectionName = "users"
const FilmsCollectionName = "films"
const CoockiesCollectionName = "usersAuthCookies"

const APIPort = ":8000"

const CertFile = "/ssl/cert.pem"
const KeyFile = "/ssl/privkey.pem"
