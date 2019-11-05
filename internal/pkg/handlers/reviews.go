package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerProfileReviews(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Request().Cookie(configs.Default.CookieAuthName)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "no cookie")
		}
		user, ok := conn.FindUserByCookie(session)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "no user session in db")
		}
		var offset, limit int
		offset, err = strconv.Atoi(ctx.Param("offset"))
		if err != nil {
			offset = 0
		}
		limit, err = strconv.Atoi(ctx.Param("limit"))
		if err != nil {
			limit = 10
		}
		reviews, _ := conn.GetReviewsOfAuthorSortedByDate(user.ID, limit, offset)
		var jsonResponse []byte
		jsonResponse = append(jsonResponse, []byte("[")...)
		for i, review := range reviews {
			jsonModel, err := review.MarshalJSON()
			if err != nil {
				continue
			}
			if i > 0 {
				jsonResponse = append(jsonResponse, []byte(",")...)
			}
			jsonResponse = append(jsonResponse, jsonModel...)
		}
		jsonResponse = append(jsonResponse, []byte("]")...)
		_, err = ctx.Response().Write(jsonResponse)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func GetHandlerReviews(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var filmID, offset, limit int
		filmID, err := strconv.Atoi(ctx.Param("film_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "specify film id")
		}
		offset, err = strconv.Atoi(ctx.QueryParam("offset"))
		if err != nil {
			offset = 0
		}
		limit, err = strconv.Atoi(ctx.QueryParam("limit"))
		if err != nil {
			limit = 10
		}
		reviews, _ := conn.GetReviewsOfFilmSortedByDate(models.ID(filmID), limit, offset)
		var jsonResponse []byte
		jsonResponse = append(jsonResponse, []byte("[")...)
		for i, review := range reviews {
			jsonModel, err := review.MarshalJSON()
			if err != nil {
				continue
			}
			if i > 0 {
				jsonResponse = append(jsonResponse, []byte(",")...)
			}
			jsonResponse = append(jsonResponse, jsonModel...)
		}
		jsonResponse = append(jsonResponse, []byte("]")...)
		_, err = ctx.Response().Write(jsonResponse)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func GetHandlerReviewsCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := common.CheckPOSTRequest(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		user, ok := conn.FindUserByCookie(session)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "no user session in db")
		}
		userID := user.ID
		rawBody, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		defer ctx.Request().Body.Close()
		newReview := models.NewReview{}
		err = newReview.UnmarshalJSON(rawBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		_, err = govalidator.ValidateStruct(newReview)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		newReview.AuthorID = userID
		e := conn.Upsert(newReview)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
