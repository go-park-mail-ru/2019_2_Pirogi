package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerPages(usecase usecase.PagesUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		body, err := usecase.GetIndexPageJSONBlob()
		if err != nil {
			return err.HTTP()
		}
		network.WriteJSONToResponse(ctx, 200, body)
		return nil
	}
}
