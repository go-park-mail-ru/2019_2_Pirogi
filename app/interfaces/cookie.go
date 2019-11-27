package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
	"net/http"
)

type cookieRepository struct {
	conn database.Database
}

func NewCookieRepository(conn database.Database) *cookieRepository {
	return &cookieRepository{
		conn: conn,
	}
}

func (c *cookieRepository) Insert(cookie model.Cookie) *model.Error {
	return c.conn.Upsert(cookie)
}

func (c *cookieRepository) Update(cookie model.Cookie) *model.Error {
	return c.conn.Upsert(cookie)
}

func (c *cookieRepository) Delete(cookie model.Cookie) *model.Error {
	return c.conn.Delete(cookie)
}

func (c *cookieRepository) Get(id model.ID) (model.Cookie, *model.Error) {
	cookieInterface, e := c.conn.Get(id, configs.Default.CookieTargetName)
	if e != nil {
		return model.Cookie{}, e
	}
	if cookie, ok := cookieInterface.(model.Cookie); !ok {
		return model.Cookie{}, model.NewError(500, "can not cast cookie")
	} else {
		return cookie, nil
	}
}

func (c *cookieRepository) GetCookieFromRequest(r *http.Request, name string) (model.Cookie, *model.Error) {
	cookieCommon, err := r.Cookie(name)
	if err != nil {
		return model.Cookie{}, model.NewError(400, err.Error())
	}
	var cookie model.Cookie
	cookie.CopyFromCommon(cookieCommon)
	return cookie, nil
}

func (c *cookieRepository) SetOnResponse(res *echo.Response, r *model.Cookie) {
	http.SetCookie(res, r.Cookie)
}

func (c *cookieRepository) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	cookie, err := network.GetCookieFromContext(ctx, configs.Default.CookieAuthName)
	if err != nil {
		return model.User{}, err
	}
	return c.conn.FindUserByCookie(cookie.Cookie)
}

func (c *cookieRepository) GetUserByCookieValue(cookieValue string) (model.User, *model.Error) {
	cookie := &http.Cookie{
		Name:     configs.Default.CookieAuthName,
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	return c.conn.FindUserByCookie(cookie)
}
