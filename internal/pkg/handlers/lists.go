package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetHandlerList(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var jsonBody []byte
		for i := 0; i < 10; i++ {
			obj, e := conn.Get(models.ID(i), "film")
			if e != nil {
				continue
			}
			film := obj.(models.Film)
			jsonModel, err := film.MarshalJSON()
			if err != nil {
				continue
			}
			jsonBody = common.UnionJSONAsArray(jsonBody, jsonModel)
		}

		_, err := ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}
