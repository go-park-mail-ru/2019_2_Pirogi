package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/review"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/user"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/infrastructure/database"
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
		var items [][]byte
		for _, review := range reviews {
			jsonModel, err := review.MarshalJSON()
			if err != nil {
				continue
			}
			items = append(items, jsonModel)
		}
		jsonResponse := common.MakeJSONArray(items)
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
		reviews, _ := conn.GetReviewsOfFilmSortedByDate(domains.ID(filmID), limit, offset)
		var items [][]byte
		for _, review := range reviews {
			reviewUser, err := conn.Get(review.AuthorID, "user")
			reviewUserTrunc := makers.MakeUserTrunc(reviewUser.(user.User))
			if err != nil {
				continue
			}
			reviewFull := makers.MakeReviewFull(review, reviewUserTrunc, domains.Mark(rand.Float32()*5))
			jsonModel, e := reviewFull.MarshalJSON()
			if e != nil {
				continue
			}
			items = append(items, jsonModel)
		}
		jsonResponse := common.MakeJSONArray(items)
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
		newReview := review.ReviewNew{}
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
