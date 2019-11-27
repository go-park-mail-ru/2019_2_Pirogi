package middleware

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/metrics"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/security"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
)

func setDefaultHeaders(w http.ResponseWriter) {
	for k, v := range configs.Headers.HeadersMap {
		w.Header().Set(k, v)
	}
}

func HeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		setDefaultHeaders(ctx.Response())
		return next(ctx)
	}
}

func PostCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Request().Method == http.MethodPost && !security.CheckNoCSRF(ctx) {
			return errors.New("invalid csrf token")
		}
		return next(ctx)
	}
}

func GetAccessLogMiddleware(logger *zap.Logger) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("id", id),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
			}
			n := res.Status
			switch {
			case n >= 500:
				logger.Error("Server error", fields...)
			case n >= 400:
				logger.Warn("Client error", fields...)
			case n >= 300:
				logger.Info("Redirection", fields...)
			default:
				logger.Info("Success", fields...)
			}
			return next(c)
		}
	}
}

func SetCSRFCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := ctx.Request().Cookie(configs.Default.CSRFCookieName)
		if err == nil {
			return next(ctx)
		}
		csrfCookie := &http.Cookie{
			Name:     configs.Default.CSRFCookieName,
			Value:    hash.SHA1(time.Now().String() + ctx.Request().RemoteAddr),
			Path:     "/",
			Expires:  time.Now().Add(configs.Default.CookieAuthDurationHours * time.Hour),
			HttpOnly: false,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(ctx.Response(), csrfCookie)
		return next(ctx)
	}
}

func CheckStatusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		metrics.ApiMetrics.IncHitsTotal()
		metrics.ApiMetrics.IncHitOfResponse(ctx.Response().Status, ctx.Request().Method, ctx.Path())
		return next(ctx)
	}
}
