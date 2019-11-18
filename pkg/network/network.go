package network

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/usecases/search"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/security"
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

func MapQueryParams(ctx echo.Context) (queryParams search.QuerySearchParams) {
	queryParams.Limit = configs.Default.DefaultEntriesLimit // limit must be positive, default value(0) is not suitable
	p := reflect.ValueOf(&queryParams).Elem()
	t := reflect.TypeOf(queryParams)
	for i := 0; i < p.NumField(); i++ {
		switch p.Field(i).Kind() {
		case reflect.Int:
			val, err := strconv.Atoi(ctx.QueryParam(strings.ToLower(t.Field(i).Name)))
			if err != nil {
				continue
			}
			p.Field(i).SetInt(int64(val))
			continue
		case reflect.String:
			param := ctx.QueryParam(strings.ToLower(t.Field(i).Name))
			if param != "" {
				p.Field(i).SetString(param)
			}
		case reflect.Slice:
			switch t.Field(i).Type.Elem().Kind() {
			case reflect.String:
				querySlice := strings.Split(ctx.QueryParam(strings.ToLower(t.Field(i).Name)), ",")
				newStringSlice := reflect.MakeSlice(reflect.TypeOf([]string{}), 0, 0)
				for _, item := range querySlice {
					if item != "" {
						newStringSlice = reflect.Append(newStringSlice, reflect.ValueOf(item))
					}
				}
				p.Field(i).Set(newStringSlice)
			case reflect.Int:
				querySlice := strings.Split(ctx.QueryParam(strings.ToLower(t.Field(i).Name)), ",")
				println(strings.ToLower(t.Field(i).Name))
				var newIntValues []int
				for _, item := range querySlice {
					value, err := strconv.Atoi(item)
					if err != nil {
						continue
					}
					newIntValues = append(newIntValues, value)
				}
				newIntSlice := reflect.MakeSlice(reflect.TypeOf([]int{}), 0, 0)
				for _, item := range newIntValues {
					newIntSlice = reflect.Append(newIntSlice, reflect.ValueOf(item))
				}
				p.Field(i).Set(newIntSlice)
			}
		}
	}
	return queryParams
}

func CheckPOSTRequest(ctx echo.Context) (session *http.Cookie, err error) {
	if !security.CheckNoCSRF(ctx) {
		return nil, errors.New("invalid CSRF")
	}
	return ctx.Request().Cookie(configs.Default.CookieAuthName)
}


func NormalizePath(path string) {
	if path[len(path)-1] != '/' {
		path += "/"
	}
}