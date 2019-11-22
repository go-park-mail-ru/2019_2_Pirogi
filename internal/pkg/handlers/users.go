package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"

	"github.com/asaskevich/govalidator"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerUsers(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, ok := auth.GetUserByRequest(ctx.Request(), conn)
		if !ok {
			return echo.NewHTTPError(401, "no auth")
		}
		rawUser, err := user.MarshalJSON()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		_, err = ctx.Response().Write(rawUser)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func GetHandlerUser(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
		}
		obj, e := conn.Get(models.ID(id), configs.Default.UserTargetName)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		user := obj.(models.User)
		userInfo := user.UserTrunc
		jsonBody, err := userInfo.MarshalJSON()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		_, err = ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func GetHandlerUsersCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if !security.CheckNoCSRF(ctx) {
			return echo.NewHTTPError(http.StatusBadRequest, "no or invalid CSRF header")
		}
		_, err := ctx.Request().Cookie(configs.Default.CookieAuthName)
		if err == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "already logged in")
		}
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer ctx.Request().Body.Close()
		newUser := models.NewUser{}
		err = newUser.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		_, err = govalidator.ValidateStruct(newUser)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		e := conn.Upsert(newUser)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		e = auth.Login(ctx, conn, newUser.Email, newUser.Password)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}

func GetHandlerUsersUpdate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer ctx.Request().Body.Close()
		updateUser := &models.UpdateUser{}
		err = updateUser.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		_, err = govalidator.ValidateStruct(updateUser)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		session, err := ctx.Request().Cookie(configs.Default.CookieAuthName)
		if err != nil {
			return echo.NewHTTPError(401, err.Error())
		}
		user, ok := conn.FindUserByCookie(session)
		if !ok {
			return echo.NewHTTPError(401, "no user with the cookie")
		}
		switch {
		case updateUser.Username != "":
			user.Username = updateUser.Username
			fallthrough
		case updateUser.Description != "":
			user.Description = updateUser.Description
		}
		e := conn.Upsert(user)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
