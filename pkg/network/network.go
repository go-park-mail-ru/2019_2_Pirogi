package network

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"io/ioutil"
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
		return -1, model.NewError(400, err.Error())
	}
	return id, nil
}

func WriteJSON(ctx echo.Context, status int, body []byte) *model.Error {
	err := ctx.JSON(status, body)
	if err != nil {
		return model.NewError(500, err.Error())
	}
	return nil
}
