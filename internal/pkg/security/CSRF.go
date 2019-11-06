package security

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
)

func CheckNoCSRF(ctx echo.Context) bool {
	tokenHeader := ctx.Request().Header.Get(configs.Default.CSRFHeader)
	if tokenHeader == "" {
		return false
	}

	cookie, err := ctx.Request().Cookie(configs.Default.CSRFCookieName)
	if err != nil {
		return false
	}
	println(tokenHeader)
	println(configs.Default.CSRFCookieName + "=" + cookie.Value)
	return tokenHeader == configs.Default.CSRFCookieName+"="+cookie.Value
}
