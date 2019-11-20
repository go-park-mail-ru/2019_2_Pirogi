package repository

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"net/http"
)

type CookieRepository interface {
	Insert(cookie model.Cookie) (model.ID, error)
	Update(cookie model.Cookie) error
	Delete(id model.ID) bool
	Get(id model.ID) model.Cookie
	GetFromRequest(r *http.Request, name string) (*model.Cookie, error)
	SetOnResponse(res *echo.Response, r *model.Cookie)
	Find(cookie *model.Cookie) model.ID
}
