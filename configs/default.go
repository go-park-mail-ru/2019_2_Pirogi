package configs

import "time"

const CookieAuthName = "cinsear_session"
const CookieAuthDuration = 10 * time.Hour

// TODO: изменить путь для пользовательских картинок
const ImageUploadPath = "/var/www/media/images/users/"
const AccessLogPath = "../../"
