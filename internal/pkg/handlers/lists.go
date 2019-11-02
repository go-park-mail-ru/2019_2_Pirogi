package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
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

		_, err := ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}
