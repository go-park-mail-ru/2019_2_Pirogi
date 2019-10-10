package handlers

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func HTTPErrorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	message := "internal server error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}
	ctx.Logger().Error(ctx.Request().URL, code, err)
	file, err := getErrorLogFile()
	if err != nil {
		ctx.Logger().Warn(err.Error())
	} else {
		defer func() {
			err := file.Close()
			if err != nil {
				ctx.Logger().Warn("can not close log file")
			}
		}()
	}

	e := models.Error{
		Status: code,
		Error:  message,
	}
	jsonError, _ := e.MarshalJSON()
	ctx.Response().WriteHeader(code)
	_, _ = fmt.Fprint(ctx.Response(), string(jsonError))
}

func getErrorLogFile() (*os.File, error) {
	if f, e := os.OpenFile(configs.ErrorLog, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); e != nil {
		return nil, e
	} else {
		return f, nil
	}
}
