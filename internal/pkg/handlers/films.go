package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerFilm(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("film_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		obj, e := conn.Get(models.ID(id), "film")
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		film := obj.(models.Film)
		jsonBody, err := film.MarshalJSON()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		_, err = ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func GetHandlerFilmCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer ctx.Request().Body.Close()
		newFilm := models.NewFilm{}
		err = newFilm.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		e := conn.InsertOrUpdate(newFilm)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
