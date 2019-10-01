package configs

import "time"

const CookieAuthName = "cinsear_session"
const CookieAuthDuration = 10 * time.Hour

// TODO: изменить путь для пользовательских картинок
//const UsersImageUploadPath = "../../media/users/"

const AccessLog = "/var/log/cinsear/access.log"
const ErrorLog = "/var/log/cinsear/error.log"

const UsersImageUploadPath = "/var/www/media/images/users/"
const FilmsImageUploadPath = "/var/www/media/images/films/"
