package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/review"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func GetHandlerReviewsCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session, err := ctx.Request().Cookie(configs.Default.CookieAuthName)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "no cookie")
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
		newReview.AuthorID = userID
		result, err := review.CreateReview(newReview)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		e := conn.InsertOrUpdate(result)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
