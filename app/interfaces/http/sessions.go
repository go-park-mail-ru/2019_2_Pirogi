package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerLoginCheck(usecase usecase.AuthUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		newEventsNumber, ok := usecase.LoginCheck(ctx)
		if !ok {
			return echo.NewHTTPError(401, "Пользователь не авторизован")
		}
		newEventsResponse := map[string]int{"new_events_number": newEventsNumber}
		jsonBlob, err := json.Marshal(newEventsResponse)
		if err != nil {
			return model.NewError(500, err.Error()).HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
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
		newEventsNumber, err := usecase.Login(ctx, credentials.Email, hash.SHA1(credentials.Password))
		if err != nil {
			return err.HTTP()
		}
		body, e := json.Marshal(map[string]int{"new_events_number": newEventsNumber})
		if e != nil {
			return model.NewError(500, e.Error()).HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, body)
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
