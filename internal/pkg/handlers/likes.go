package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/labstack/echo"
)

func GetHandlerLikesCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ok := security.CheckNoCSRF(ctx)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "No CSRF token")
		}
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
		newLike := domains.Like{}
		err = newLike.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		_, err = govalidator.ValidateStruct(newLike)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		newLike.UserID = userID
		e := conn.Upsert(newLike)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
