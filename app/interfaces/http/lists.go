package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/labstack/echo"
)

func GetHandlerLists(listsUsecase usecase.ListsUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {

	}
}


func GetHandlerCreateOrUpdate(listsUsecase usecase.ListsUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {

	}
}
