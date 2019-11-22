package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerPerson(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("person_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
		}
		obj, e := conn.Get(models.ID(id), "person")
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		person := obj.(models.Person)
		films, _ := conn.FindFilmsByIDs(person.FilmsID)
		var filmPersonsTrunc [][]models.PersonTrunc
		for _, f := range films {
			persons, _ := conn.FindPersonsByIDs(f.PersonsID)
			personsTrunc := makers.MakePersonsTrunc(persons)
			filmPersonsTrunc = append(filmPersonsTrunc, personsTrunc)
		}
		personFull := makers.MakeFullPerson(person, films, filmPersonsTrunc)
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
		model, err := common.PrepareModel(rawBody, models.NewPerson{})
		newPerson := model.(models.NewPerson)
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
