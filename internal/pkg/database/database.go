package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"net/http"
)

type Database interface {
	InsertOrUpdate(in interface{}) *models.Error
	Get(id models.ID, target string) (interface{}, *models.Error)
	Delete(in interface{}) *models.Error
	CheckCookie(cookie *http.Cookie) bool
	FindUserByEmail(email string) (models.User, bool)
	FindUserByID(id models.ID) (models.User, bool)
	FindUserByCookie(cookie *http.Cookie) (models.User, bool)
	FindFilmByTitle(title string) (models.Film, bool)
	FindFilmByID(id models.ID) (models.Film, bool)
	ClearDB()
}
