
package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetHandlerList(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		queryParams := common.MapQueryParams(ctx)
		films, e := common.GetByQueryListParams(conn, queryParams)
		var filmPersonsTrunc [][]models.PersonTrunc
		for _, f := range films {
			persons, _ := conn.FindPersonsByIDs(f.PersonsID)
			personsTrunc := makers.MakePersonsTrunc(persons)
			filmPersonsTrunc = append(filmPersonsTrunc, personsTrunc)
		}
		filmsTrunc := makers.MakeFilmsTrunc(films, filmPersonsTrunc)
		if e != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, e)
		}
		err := ctx.JSON(200, filmsTrunc)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}