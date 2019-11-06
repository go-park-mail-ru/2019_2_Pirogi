package common

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/security"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ReadBody(ctx echo.Context) ([]byte, error) {
	rawBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}
	err = ctx.Request().Body.Close()
	if err != nil {
		return nil, err
	}
	return rawBody, nil
}

func PrepareModel(body []byte, in interface{}) (out interface{}, err error) {
	switch in.(type) {
	case models.NewFilm:
		newModel := in.(models.NewFilm)
		err = newModel.UnmarshalJSON(body)
		if err != nil {
			return nil, err
		}
		_, err = govalidator.ValidateStruct(newModel)
		return newModel, err
	case models.NewReview:
		newModel := in.(models.NewReview)
		err = newModel.UnmarshalJSON(body)
		if err != nil {
			return nil, err
		}
		_, err = govalidator.ValidateStruct(newModel)
		return newModel, err
	case models.NewUser:
		newModel := in.(models.NewUser)
		err = newModel.UnmarshalJSON(body)
		if err != nil {
			return nil, err
		}
		_, err = govalidator.ValidateStruct(newModel)
		return newModel, err
	case models.NewPerson:
		newModel := in.(models.NewPerson)
		err = newModel.UnmarshalJSON(body)
		if err != nil {
			return nil, err
		}
		_, err = govalidator.ValidateStruct(newModel)
		return newModel, err
	default:
		return nil, errors.New("unsupported model")
	}
}

func CheckPOSTRequest(ctx echo.Context) (session *http.Cookie, err error) {
	if !security.CheckNoCSRF(ctx) {
		return nil, errors.New("invalid CSRF")
	}
	return ctx.Request().Cookie(configs.Default.CookieAuthName)
}

func MapQueryListParams(ctx echo.Context) (queryParams models.QueryListParams) {
	queryParams.Limit = configs.Default.DefaultEntriesLimit // limit must be positive, default value(0) is not suitable
	p := reflect.ValueOf(&queryParams).Elem()
	t := p.Type()
	for i := 0; i < p.NumField(); i++ {
		if p.Field(i).Kind() == reflect.Int {
			val, err := strconv.Atoi(ctx.QueryParam(strings.ToLower(t.Field(i).Name)))
			if err != nil {
				continue
			}
			p.Field(i).SetInt(int64(val))
			continue
		}
		p.Field(i).SetString(ctx.QueryParam(strings.ToLower(t.Field(i).Name)))
	}
	return
}

func GetByQueryListParams(conn database.Database, qp models.QueryListParams) ([]models.Film, *models.Error) {
	var (
		items []models.Film
		err   *models.Error
	)
	if qp.Genre != "" {
		items, err = conn.GetFilmsOfGenreSortedByMark(models.Genre(qp.Genre), qp.Limit, qp.Offset)
	} else {
		items, err = conn.GetFilmsSortedByMark(qp.Limit, qp.Offset)
	}
	return items, err
}
