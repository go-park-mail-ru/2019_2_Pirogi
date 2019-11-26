package network

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ReadBody(ctx echo.Context) ([]byte, *model.Error) {
	rawBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, &model.Error{
			Status: 400,
			Error:  err.Error(),
		}
	}
	err = ctx.Request().Body.Close()
	if err != nil {
		return nil, &model.Error{
			Status: 400,
			Error:  err.Error(),
		}
	}
	return rawBody, nil
}

func NormalizePath(path string) {
	if path[len(path)-1] != '/' {
		path += "/"
	}
}

func GetIntParam(ctx echo.Context, param string) (int, *model.Error) {
	id, err := strconv.Atoi(ctx.Param(param))
	if err != nil {
		return -1, model.NewError(400, param, err.Error())
	}
	return id, nil
}

func WriteJSONToResponse(ctx echo.Context, status int, body []byte) {
	err := ctx.JSONBlob(status, body)
	if err != nil {
		zap.S().Warn(err.Error())
	}
}

func GetCookieFromContext(ctx echo.Context, name string) (model.Cookie, *model.Error) {
	cookieCommon, err := ctx.Request().Cookie(name)
	if err != nil {
		return model.Cookie{}, model.NewError(400, err.Error())
	}
	var cookie model.Cookie
	cookie.CopyFromCommon(cookieCommon)
	return cookie, nil
}

func SetCookieOnContext(ctx *echo.Context, cookie model.Cookie) {
	http.SetCookie((*ctx).Response(), cookie.Cookie)
}
