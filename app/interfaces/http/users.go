package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"net/http"

	"github.com/labstack/echo"
)

func GetHandlerUsers(userUsecase usecase.UserUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := userUsecase.GetUserByContext(ctx)
		if err != nil {
			return echo.NewHTTPError(401, "no auth")
		}
		user.Password = ""
		jsonBody, e := user.MarshalJSON()
		if e != nil {
			return echo.NewHTTPError(400, "invalid request data")
		}
		network.WriteJSONToResponse(ctx, 200, jsonBody)
		return nil
	}
}

func GetHandlerUser(userUsecase usecase.UserUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := network.GetIntParam(ctx, "user_id")
		if err != nil {
			return err.HTTP()
		}
		user, err := userUsecase.GetUserTruncByteByID(model.ID(id))
		if err != nil {
			return err.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, user)
		return nil
	}
}

func GetHandlerUsersCreate(userUsecase usecase.UserUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, _ := userUsecase.GetUserByContext(ctx)
		if user.Email != "" {
			return echo.NewHTTPError(http.StatusBadRequest, "already logged in")
		}
		err := userUsecase.CreateUserNewFromContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}

func GetHandlerUsersUpdate(userUsecase usecase.UserUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := userUsecase.GetUserByContext(ctx)
		if err != nil {
			return echo.NewHTTPError(401, "no auth")
		}
		err = userUsecase.UpdateUserFromContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
