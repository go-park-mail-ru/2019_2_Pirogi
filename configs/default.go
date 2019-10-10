package configs

import "time"

const CookieAuthName = "cinsear_session"
const CookieAuthDuration = 10 * time.Hour

// TODO: изменить путь для пользовательских картинок
// const UsersImageUploadPath = "../../media/users/"

const AccessLog = "/cinsear/log/access.log"
const ErrorLog = "/cinsear/log/error.log"

const UsersImageUploadPath = "/var/www/media/images/users/"
const FilmsImageUploadPath = "/var/www/media/images/films/"

const MongoDbUri = "mongodb://127.0.0.1:27017"
const MongoDbName = "cinsear"
const UsersCollectionName = "users"
const FilmsCollectionName = "films"
const CoockiesCollectionName = "usersAuthCookies"

const APIPort = ":8000"
const CertFile = "/cinsear/ssl/cert.pem"
const KeyFile = "/cinsear/ssl/privkey.pem"
