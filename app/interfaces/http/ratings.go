package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerRatingsCreateOrUpdate(u usecase.RatingUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := u.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		err = u.CreateOrUpdateRating(rawBody, user)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
