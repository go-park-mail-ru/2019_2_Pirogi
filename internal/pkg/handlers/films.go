package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
)

func GetHandlerFilm(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("film_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		obj, e := conn.Get(models.ID(id), configs.Default.FilmTargetName)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		film := obj.(models.Film)
		persons, _ := conn.FindPersonsByIDs(film.PersonsID)
		filmFull := makers.MakeFilmFull(film, persons)
		jsonBody, err := filmFull.MarshalJSON()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		_, err = ctx.Response().Write(jsonBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func GetHandlerFilmCreate(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := common.CheckPOSTRequest(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		rawBody, err := common.ReadBody(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		model, err := common.PrepareModel(rawBody, models.NewFilm{})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		newFilm, _ := model.(models.NewFilm)
		//TODO: было бы классно, если он возвращал ID
		e := conn.Upsert(newFilm)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		film, ok := conn.FindFilmByTitle(newFilm.Title)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Upsert failed")
		}

		//TODO: сделать это ассинхронным и красивым, потому что сейчас вообще г, даже смотреть противно, был бы этот код собакой, я бы усыпил
		persons, _ := conn.FindPersonsByIDs(film.PersonsID)
		for idx, person := range persons {
			flag := true
			for _, filmID := range person.FilmsID {
				if filmID == film.ID {
					flag = false
					break
				}
			}
			if flag {
				persons[idx].FilmsID = append(person.FilmsID, film.ID)
			}

			for _, filmGenre := range film.Genres {
				flag = true
				for _, personGenre := range person.Genres {
					if filmGenre == personGenre {
						flag = false
						break
					}
				}
				if flag {
					persons[idx].Genres = append(person.Genres, filmGenre)
				}
			}
			conn.Upsert(persons[idx])
		}
		return nil
	}
}
