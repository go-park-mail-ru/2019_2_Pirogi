package database

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

type Database interface {
	Insert(in interface{}) *models.Error
	Get(id int, target string) (interface{}, *models.Error)
	Delete(in interface{}) *models.Error
	CheckCookie(cookie *http.Cookie) bool
	FindUserByEmail(email string) (models.User, bool)
	FindUserByID(id int) (models.User, bool)
	FindUserByCookie(cookie *http.Cookie) (models.User, bool)
	FindFilmByTitle(title string) (models.Film, bool)
	FindFilmByID(id int) (models.Film, bool)
	FakeFillDB()
}
