package http

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerProfileReviews(u usecase.ReviewUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := u.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		limit, offset := u.GetLimitAndOffset(ctx)
		jsonBlob, err := u.GetUserReviewsJSONBlob(user, limit, offset)
		if err != nil {
			return err.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
		return nil
	}
}

func GetHandlerReviews(u usecase.ReviewUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		filmID, err := strconv.Atoi(ctx.Param("film_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "specify film id")
		}
		limit, offset := u.GetLimitAndOffset(ctx)
		jsonBlob, e := u.GetFilmReviewsFullJSONBlob(model.ID(filmID), limit, offset)
		if e != nil {
			return e.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
		return nil
	}
}

func GetHandlerReviewsCreate(u usecase.ReviewUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := u.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		err = u.CreateNewReview(rawBody, user)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
