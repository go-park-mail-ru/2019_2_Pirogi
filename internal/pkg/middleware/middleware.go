package middleware

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"time"
)

func setDefaultHeaders(w http.ResponseWriter) {
	for k, v := range configs.Headers {
		w.Header().Set(k, v)
	}
}

func HeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		setDefaultHeaders(ctx.Response())
		return next(ctx)
	}
}

func AccessLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if f, err := os.OpenFile(configs.AccessLog, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
			ctx.Logger().Warn(err.Error())
		} else {
			_, err = fmt.Fprintf(f, "%s %s %s %s \n", time.Now().Format("02/01 15:04:05"),
				ctx.Request().Method, ctx.Request().URL, ctx.Request().Host)
			if err != nil {
				ctx.Logger().Warn(err.Error())
			}
		}
		return next(ctx)
	}
}
