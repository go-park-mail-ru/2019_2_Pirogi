package common

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth/security"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
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
		return nil, errors.New("no CSRF cookie")
	}
	return ctx.Request().Cookie(configs.Default.CookieAuthName)
}
