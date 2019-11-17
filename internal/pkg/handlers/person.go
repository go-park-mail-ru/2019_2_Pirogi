package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/labstack/echo"
)

func GetHandlerPerson(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("person_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
		}
		obj, e := conn.Get(domains.ID(id), "person")
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		person := obj.(domains.Person)
		films, _ := conn.FindFilmsByIDs(person.FilmsID)
		personFull := makers.MakeFullPerson(person, films)
		jsonBody, err := personFull.MarshalJSON()
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

func GetHandlerPersonsCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := common.CheckPOSTRequest(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		rawBody, err := common.ReadBody(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		model, err := common.PrepareModel(rawBody, domains.NewPerson{})
		newPerson := model.(domains.NewPerson)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		e := conn.Upsert(newPerson)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
