package handlers

import (
	"github.com/asaskevich/govalidator"
	"io/ioutil"
	"net/http"
	"strconv"

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
		jsonBody, err := person.MarshalJSON()
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
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer ctx.Request().Body.Close()
		newPerson := models.NewPerson{}
		err = newPerson.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		_, err = govalidator.ValidateStruct(newPerson)
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
