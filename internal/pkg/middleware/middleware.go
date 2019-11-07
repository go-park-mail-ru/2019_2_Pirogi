package middleware

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"regexp"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/labstack/echo"
)

func ExpireInvalidCookiesMiddleware(conn database.Database) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, err := c.Request().Cookie(configs.Default.CookieAuthName)
			if err != nil {
				return next(c)
			}
			_, ok := conn.FindUserByCookie(session)
			if !ok {
				auth.ExpireCookie(session)
				http.SetCookie(c.Response(), session)
				return next(c)
			}
			return next(c)
		}
	}
}

func setDefaultHeaders(w http.ResponseWriter, origin string) {
	var ipWithPortRegexp = regexp.MustCompile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])/.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+$")
	for k, v := range configs.Headers.HeadersMap {
		w.Header().Set(k, v)
	}
	log.Warn(origin)
	if ipWithPortRegexp.MatchString(origin) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	} else {
		w.Header().Set("Access-Control-Allow-Origin", "https://cinsear.ru")

	}
}

func HeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		setDefaultHeaders(ctx.Response(), ctx.Request().RemoteAddr)
		return next(ctx)
	}
}

func GetAccessLogMiddleware(logger *zap.Logger) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
			}
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
			return nil
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
			Value:    user.GetMD5Hash(time.Now().String() + ctx.Request().RemoteAddr),
			Path:     "/",
			Expires:  time.Now().Add(configs.Default.CookieAuthDurationHours * time.Hour),
			HttpOnly: false,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(ctx.Response(), csrfCookie)
		return next(ctx)
	}
}
