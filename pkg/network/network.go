package network

import (
	"io/ioutil"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"go.uber.org/zap"
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
