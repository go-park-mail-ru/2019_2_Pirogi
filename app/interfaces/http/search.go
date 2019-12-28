package http

import (
	usecase2 "github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerSearch(usecase usecase2.SearchUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		jsonBody, e := usecase.GetFilmsByGetParamsJSONBlob(ctx)
		// Проверяем, что jsonBody вернулся не пустой, т.е. не []
		if e == nil && len(jsonBody) > 2 {
			network.WriteJSONToResponse(ctx, 200, jsonBody)
			return nil
		}

		jsonBody, e = usecase.GetPersonsByGetParamsJSONBlob(ctx)
		network.WriteJSONToResponse(ctx, 200, jsonBody)
		return nil
	}
}
