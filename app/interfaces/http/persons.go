package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	json2 "github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerPerson(u usecase.PersonUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var isAuth, isSubscribed bool
		id, err := network.GetIntParam(ctx, "person_id")
		if err != nil {
			return err.HTTP()
		}
		body, err := u.GetPersonFullByte(model.ID(id))
		if err != nil {
			return err.HTTP()
		}

		user, err := u.GetUserByContext(ctx)
		if err == nil {
			isAuth = true
			subscriptions := u.CheckSubscription(user.ID, model.ID(id))
			isSubscribed = subscriptions
		}
		newEventsNumber, err := u.GetNewEventsNumber(user.ID)
		if err != nil {
			newEventsNumber = 0
		}
		params, e := json.Marshal(map[string]interface{}{"is_auth": isAuth, "is_subscribed": isSubscribed,
			"new_events_number": newEventsNumber})
		if e != nil {
			return model.NewError(500, e.Error()).HTTP()
		}
		jsonBlob := json2.UnionToJSONBytes([]string{"params", "person",},
			[][]byte{params, body})
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
		return nil
	}
}

func GetHandlerPersonsCreate(u usecase.PersonUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		err = u.Create(rawBody)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
