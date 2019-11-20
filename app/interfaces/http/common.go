package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
)

func HandlerCommon() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		variable := ctx.Param("variable")
		var err error
		switch variable {
		case "genres":
			err = ctx.JSON(200, configs.Genres)
		case "roles":
			err = ctx.JSON(200, configs.Roles)
		default:
			return echo.NewHTTPError(400, "unsupported type: ", variable)
		}
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		return nil
	}
}
