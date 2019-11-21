package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	usecase2 "github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerLoginCheck(usecase usecase2.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ok := usecase.LoginCheck(ctx)
		if !ok {
			return echo.NewHTTPError(401, "no auth")
		}
		return nil
	}
}

func GetHandlerLogin(usecase usecase2.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		credentials := model.UserCredentials{}
		e := credentials.UnmarshalJSON(rawBody)
		if e != nil {
			return e
		}
		err = usecase.Login(ctx, credentials.Email, credentials.Password)
		if err != nil {
			return err.Common()
		}
		return nil
	}
}

func GetHandlerLogout(usecase usecase2.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		e := usecase.Logout(ctx)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}
