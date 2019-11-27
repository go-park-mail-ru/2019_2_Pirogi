package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerSubscribe(usecase usecase.SubscriptionUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := usecase.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		body, err := network.ReadBody(ctx)
		var request model.SubscriptionRequest
		e := request.UnmarshalJSON(body)
		if e != nil {
			return model.NewError(400, e.Error()).HTTP()
		}
		err = usecase.Subscribe(u.ID, request.PersonID)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}

func GetHandlerUnsubscribe(usecase usecase.SubscriptionUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := usecase.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		body, err := network.ReadBody(ctx)
		var request model.SubscriptionRequest
		e := request.UnmarshalJSON(body)
		if e != nil {
			return model.NewError(400, e.Error()).HTTP()
		}
		err = usecase.Unsubscribe(u.ID, request.PersonID)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}

func GetHandlerNewEvents(usecase usecase.SubscriptionUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := usecase.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		jsonBlob, err := usecase.GetNewEventsListJSONBlob(u.ID)
		if err != nil {
			return err.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
		return nil
	}
}

func GetHandlerSubscriptionList(usecase usecase.SubscriptionUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := usecase.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		jsonBlob, err := usecase.GetPersonsTruncListJSONBlob(u.ID)
		if err != nil {
			return err.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
		return nil
	}
}

func GetHandlerReadNewEvents(usecase usecase.SubscriptionUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, err := usecase.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		err = usecase.ReadAllNewEvents(u.ID)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
