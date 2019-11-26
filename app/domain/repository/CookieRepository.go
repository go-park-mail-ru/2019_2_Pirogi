package repository

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"net/http"
)

type CookieRepository interface {
	Insert(cookie model.Cookie) *model.Error
	Update(cookie model.Cookie) *model.Error
	Delete(cookie model.Cookie) *model.Error
	Get(id model.ID) (model.Cookie, *model.Error)
	GetCookieFromRequest(r *http.Request, name string) (model.Cookie, *model.Error)
	SetOnResponse(res *echo.Response, r *model.Cookie)
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
}
