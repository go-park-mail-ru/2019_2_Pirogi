package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerLists(listsUsecase usecase.ListsUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		lists, err := listsUsecase.GetListsByUserCtx(ctx)
		if err != nil {
			return err.HTTP()
		}
		body := modelWorker.MarshalListsFull(lists)
		jsonBlob := json.UnionToJSONBytes([]string{"lists"}, body)
		network.WriteJSONToResponse(ctx, 200, jsonBlob)
		return nil
	}
}

func GetHandlerCreateOrUpdateList(listsUsecase usecase.ListsUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := listsUsecase.CreateOrUpdateList(ctx)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
