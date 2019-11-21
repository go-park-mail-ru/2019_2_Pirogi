package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	usecase "github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase/person"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerPerson(u usecase.PersonUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := network.GetIntParam(ctx, "person_id")
		if err != nil {
			return err.HTTP()
		}
		body, err := u.GetPersonFullByte(model.ID(id))
		if err != nil {
			return err.HTTP()
		}
		err = network.WriteJSONToResponse(ctx, 200, body)
		if err != nil {
			return err.HTTP()
		}
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
