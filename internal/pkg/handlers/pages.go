package handlers

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/search"
	"github.com/labstack/echo"
)

func GetHandlerPages(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		films, e := common.GetByQueryListParams(conn, search.QuerySearchParams{
			Limit: configs.Default.DefaultEntriesLimit + 5,
		})
		trailers := makers.MakeTrailersList(films)
		var filmPersonsTrunc [][]models.PersonTrunc
		for _, f := range films {
			persons, _ := conn.FindPersonsByIDs(f.PersonsID)
			personsTrunc := makers.MakePersonsTrunc(persons)
			filmPersonsTrunc = append(filmPersonsTrunc, personsTrunc)
		}
		filmsTrunc := makers.MakeFilmsTrunc(films, filmPersonsTrunc)
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
