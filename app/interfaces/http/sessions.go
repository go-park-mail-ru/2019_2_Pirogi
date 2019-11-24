package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerLoginCheck(usecase usecase.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ok := usecase.LoginCheck(ctx)
		if !ok {
			return echo.NewHTTPError(401, "Пользователь не авторизован")
		}
		return nil
	}
}

func GetHandlerLogin(usecase usecase.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		credentials := model.UserCredentials{}
		e := credentials.UnmarshalJSON(rawBody)
		if e != nil {
			return model.NewError(400, "Невалидные входные данные").HTTP()
		}
		err = usecase.Login(ctx, credentials.Email, credentials.Password)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}

func GetHandlerLogout(usecase usecase.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		e := usecase.Logout(ctx)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
