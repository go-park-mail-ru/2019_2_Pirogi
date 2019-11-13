package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func GetHandlerSearch(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		qp := common.MapQueryParams(ctx)
		zap.S().Debug("SearchHandler: Query params: ", qp)
		films, e := conn.GetByQuery(configs.Default.FilmsCollectionName, qp.GetPipelineForMongo("films"))
		zap.S().Debug("SearchHandler: Found films: ", films)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		err := ctx.JSON(200, films)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
		return nil
	}
}
