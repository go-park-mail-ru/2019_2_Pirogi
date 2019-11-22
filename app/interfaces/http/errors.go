package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
	"time"
)

func GetHTTPErrorHandler(logger *zap.Logger) func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		e := model.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
		if he, ok := err.(*echo.HTTPError); ok {
			e.Status = he.Code
			switch he.Message.(type) {
			case string:
				e.Error = err.Error()
			case int:
				e.Error = strconv.Itoa(he.Message.(int))
			case *model.Error:
				e.Error = he.Message.(*model.Error).Error
			}
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
