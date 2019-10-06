package configs

import "time"

const CookieAuthName = "cinsear_session"
const CookieAuthDuration = 10 * time.Hour

// TODO: изменить путь для пользовательских картинок
//const UsersImageUploadPath = "../../media/users/"

const AccessLog = "../../../access.log"
const ErrorLog = "../../../error.log"

const UsersImageUploadPath = "/var/www/media/images/users/"
const FilmsImageUploadPath = "/var/www/media/images/films/"

const CertFile = "/etc/letsencrypt/live/cinsear.online/cert.pem"
const KeyFile = "/etc/letsencrypt/live/cinsear.online/privkey.pem"
