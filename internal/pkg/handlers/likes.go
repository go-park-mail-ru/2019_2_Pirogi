package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func GetHandlerLikesCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Request().Cookie(configs.Default.CookieAuthName)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "no cookie")
		}
		user, ok := conn.FindUserByCookie(session)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "no user session in db")
		}
		userID := user.ID
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer ctx.Request().Body.Close()
		newLike := models.Like{}
		err = newLike.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		newLike.UserID = userID
		e := conn.InsertOrUpdate(newLike)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
