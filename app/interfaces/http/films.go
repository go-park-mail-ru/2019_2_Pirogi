package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	usecase2 "github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	json2 "github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

func GetHandlerFilm(filmUsecase usecase2.FilmUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := network.GetIntParam(ctx, "film_id")
		if err != nil {
			return err.HTTP()
		}
		filmID := model.ID(id)
		filmBody, err := filmUsecase.GetFilmFullByte(filmID)
		if err != nil {
			return err.HTTP()
		}
		params := map[string]interface{}{
			"is_auth":     false,
			"lists":       []string{},
			"active_list": "",
			"stars":       "",
		}
		user, err := filmUsecase.GetUserByContext(ctx)
		if err == nil {
			params["is_auth"] = true
		}
		if params["is_auth"].(bool) {
			lists, err := filmUsecase.GetUserLists(user)
			if err == nil {
				var listsTitles []string
				for _, list := range lists {
					listsTitles = append(listsTitles, list.Title)
				}
				params["lists"] = listsTitles
				params["active_list"] = filmUsecase.GetActiveList(filmID, lists)
			}
		}

		if params["is_auth"].(bool) {
			stars := filmUsecase.GetStars(filmID, user.ID)
			params["stars"] = stars
		}

		paramsBody, e := json.Marshal(params)
		if e != nil {
			return echo.NewHTTPError(500, "Ошибка при обработке параметров"+e.Error())
		}
		jsonBody := json2.UnionToJSONBytes([]string{"params", "film"}, [][]byte{paramsBody, filmBody})
		network.WriteJSONToResponse(ctx, 200, jsonBody)
		return nil
	}
}

func GetHandlerFilmCreate(filmUsecase usecase2.FilmUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rawBody, err := network.ReadBody(ctx)
		if err != nil {
			return err.HTTP()
		}
		err = filmUsecase.Create(rawBody)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
