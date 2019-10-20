package handlers

import (
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerFilm(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("film_id"))
		if err != nil {
			return echo.NewHTTPError(404, err.Error())
		}
		obj, e := conn.Get(id, "film")
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		film := obj.(models.Film)
		jsonBody, err := film.MarshalJSON()
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		_, err = ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		return nil
	}
}

func GetHandlerFilmCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}
