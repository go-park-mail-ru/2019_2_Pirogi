package common

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func ReadBody(ctx echo.Context) ([]byte, error) {
	rawBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = ctx.Request().Body.Close()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return rawBody, nil
}

func PrepareModel(body []byte, in interface{}) (out interface{}, err error) {
	switch in.(type) {
	case models.NewFilm:
		err = in.(models.NewFilm).UnmarshalJSON(body)
	case models.NewReview:
		err = in.(models.NewReview).UnmarshalJSON(body)
	case models.NewUser:
		err = in.(models.NewUser).UnmarshalJSON(body)
	case models.NewPerson:
		err = in.(models.NewPerson).UnmarshalJSON(body)
	default:
		err = echo.NewHTTPError(500, "Invalid model")
	}
	if err != nil {
		return nil, err
	}
	_, err = govalidator.ValidateStruct(in)
	if err != nil {
		return nil, err
	}
	in = FilterXSS(in)
	return in, nil
}
