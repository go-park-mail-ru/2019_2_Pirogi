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
		queryResult, e := conn.GetByQuery(configs.Default.FilmsCollectionName, qp.GetPipelineForMongo(configs.Default.FilmTargetName))
		zap.S().Debug("SearchHandler: Found films: ", queryResult)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		if len(queryResult) == 0 {
			queryResult, e = conn.GetByQuery(configs.Default.PersonsCollectionName, qp.GetPipelineForMongo(configs.Default.PersonTargetName))
			zap.S().Debug("SearchHandler: Found persons: ", queryResult)
			if e != nil {
				return echo.NewHTTPError(e.Status, e.Error)
			}
		}
		err := ctx.JSON(200, queryResult)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
		return nil
	}
}
