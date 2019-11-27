package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	usecase2 "github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerFilm(filmUsecase usecase2.FilmUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := network.GetIntParam(ctx, "film_id")
		if err != nil {
			return err.HTTP()
		}
		jsonBody, err := filmUsecase.GetFilmFullByte(model.ID(id))
		if err != nil {
			return err.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, jsonBody)
		return nil
	}
}

func GetHandlerFilmCreate(filmUsecase usecase2.FilmUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		err = filmUsecase.Create(rawBody)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
