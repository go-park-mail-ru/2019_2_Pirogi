package handlers

import (
	"io/ioutil"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerLoginCheck(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ok := auth.LoginCheck(ctx.Response(), ctx.Request(), conn)
		if !ok {
			return echo.NewHTTPError(401, "no auth")
		}
		return nil
	}
}

func GetHandlerLogin(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		defer ctx.Request().Body.Close()
		credentials := models.Credentials{}
		err = credentials.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		e := auth.Login(ctx, conn, credentials.Email, credentials.Password)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}

func GetHandlerLogout(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		e := auth.Logout(ctx, conn)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
