package handlers

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
//	"github.com/labstack/echo"
//)
//
//func GetHandlerCommon() echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		value := ctx.Param("variable")
//		switch value {
//		case "genres":
//			jsonResponse, _ := json.Marshal(struct {
//				Genres []string `json:"genres"`
//			}{Genres: configs.Genres})
//			_, _ = fmt.Fprint(ctx.Response(), jsonResponse)
//		case "roles":
//			jsonResponse, _ := json.Marshal(struct {
//				Roles []string `json:"roles"`
//			}{Roles: configs.Roles})
//			_, _ = fmt.Fprint(ctx.Response(), jsonResponse)
//		default:
//			return echo.NewHTTPError(400, "Unsupported option")
//		}
//		return nil
//	}
//}
