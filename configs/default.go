package configs

import "time"

const CookieAuthName = "cinsear_session"
const CookieAuthDuration = 10 * time.Hour

// TODO: изменить путь для пользовательских картинок
//const UsersImageUploadPath = "../../media/users/"

const AccessLog = "/log/access.log"
const ErrorLog = "/log/error.log"

const UsersImageUploadPath = "/media/images/users/"
const FilmsImageUploadPath = "/media/images/films/"

const CertFile = "../../ssl/cert.pem"
const KeyFile = "../../ssl/privkey.pem"
