package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerList(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var jsonBody []byte
		limit, err := strconv.Atoi(ctx.QueryParam("limit"))
		if limit == 0 || err != nil {
			limit = 10
		}
		for i := 0; i < limit; i++ {
			obj, e := conn.Get(models.ID(i), "film")
			if e != nil {
				continue
			}
			film := obj.(models.Film)
			persons, _ := conn.FindPersonsByIDs(film.PersonsID)
			filmFull := makers.MakeFullFilm(film, persons)
			jsonModel, err := filmFull.MarshalJSON()
			if err != nil {
				continue
			}
			if i > 0 {
				jsonBody = append(jsonBody, []byte(",")...)
			}
			jsonBody = append(jsonBody, jsonModel...)
		}

		_, err = ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}
