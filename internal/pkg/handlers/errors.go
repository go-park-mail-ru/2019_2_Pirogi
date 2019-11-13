package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
	"time"
)

func GetHTTPErrorHandler(logger *zap.Logger) func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		e := models.Error{
			Status: http.StatusInternalServerError,
			Error:  "internal server error",
		}
		if he, ok := err.(*echo.HTTPError); ok {
			e.Status = he.Code
			switch he.Message.(type) {
			case string:
				e.Error = he.Message.(string)
			case int:
				e.Error = strconv.Itoa(he.Message.(int))
			}
		} else {
			e.Error = err.Error()
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
