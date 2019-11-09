package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetHandlerPages(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		films, e := common.GetByQueryListParams(conn, models.QuerySearchParams{
			Limit: configs.Default.DefaultEntriesLimit + 5,
		})
		trailers := makers.MakeTrailersList(films)
		filmsTrunc := makers.MakeFilmsTrunc(films)
		if e != nil {
			return echo.NewHTTPError(500, e)
		}
		err := ctx.JSONBlob(200,
			common.UnionToJSON([]string{"filmsNew", "filmsForUser", "trailers"}, filmsTrunc[:8], filmsTrunc[8:], trailers))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}
