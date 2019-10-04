package database

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

type Database interface {
	Insert(in interface{}) *models.Error
	InsertCookie(cookie http.Cookie, id int) *models.Error
	DeleteCookie(in interface{})
	Get(id int, target string) (interface{}, *models.Error)
	FindByEmail(email string) (models.User, bool)
	FindUserByID(id int) (models.User, bool)
	FindUserByCookie(cookie http.Cookie) (models.User, bool)
	FindFilmByTitle(title string) (models.Film, bool)
	FakeFillDB()
	CheckCookie(cookie http.Cookie) bool
}
