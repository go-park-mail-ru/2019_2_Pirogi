package collections

//import (
//	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
//	"github.com/labstack/echo"
//	"net/http"
//)
//
//func GetHandlerList(conn database.Database) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		queryParams := common.MapQueryParams(ctx)
//		films, e := common.GetByQueryListParams(conn, queryParams)
//		filmsTrunc := makers.MakeFilmsTrunc(films)
//		if e != nil {
//			return echo.NewHTTPError(http.StatusInternalServerError, e)
//		}
//		err := ctx.JSON(200, filmsTrunc)
//		if err != nil {
//			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
//		}
//		return nil
//	}
//}