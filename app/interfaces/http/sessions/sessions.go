package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	usecase "github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerLoginCheck(usecase usecase.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ok := usecase.LoginCheck(ctx)
		if !ok {
			return echo.NewHTTPError(401, "no auth")
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
		err = credentials.UnmarshalJSON(rawBody)
		if err != nil {
			return err.HTTP()
		}
		e := usecase.Login(ctx, credentials.Email, credentials.Password)
		if e != nil {
			return e.HTTP()
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
