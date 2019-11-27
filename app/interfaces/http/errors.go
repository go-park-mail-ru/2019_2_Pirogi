package http

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetHTTPErrorHandler(logger *zap.Logger) func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		e := model.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		if he, ok := err.(*echo.HTTPError); ok {
			e.Status = he.Code
			e.Error = he.Message.(string)
		}

		fields := []zapcore.Field{
			zap.Int("status", e.Status),
			zap.String("time", time.Now().String()),
			zap.String("message", e.Error),
		}
		logger.Error("Error: ", fields...)
		err = ctx.JSON(e.Status, e)
	}
}
