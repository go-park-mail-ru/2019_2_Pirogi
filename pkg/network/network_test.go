package network

import (
	"bufio"
	"bytes"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	actual := "/api/users"
	expected := "/api/users/"
	NormalizePath(actual)
	require.Equal(t, expected, actual)
}

func TestReadBody(t *testing.T) {
	input := []byte{'a', 'b', 'c', 'd'}
	expected := []byte{'a', 'b', 'c', 'd'}
	e := echo.New()
	reader := bufio.NewReader(bytes.NewBuffer(input))
	req := httptest.NewRequest("POST", "http://www.google.com", reader)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	actual, err := ReadBody(ctx)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestMapQueryParams(t *testing.T) {
	configPath := "../../../configs"
	err := configuration.UnmarshalConfigs(configPath)
	require.NoError(t, err)
	e := echo.New()
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://cinsear.ru", reader)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	qp := MapQueryParams(ctx)
	require.Equal(t, configs.Default.DefaultEntriesLimit, qp.Limit)
}

func TestMapQueryParamsSliceString(t *testing.T) {
	configPath := "../../../configs"
	err := configuration.UnmarshalConfigs(configPath)
	require.NoError(t, err)
	e := echo.New()
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://cinsear.ru", reader)
	q := req.URL.Query()
	q.Add("countries", "USA,Russia")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	qp := MapQueryParams(ctx)
	require.Equal(t, []string{"USA", "Russia"}, qp.Countries)
}

func TestMapQueryParamsSliceInt(t *testing.T) {
	configPath := "../../../configs"
	err := configuration.UnmarshalConfigs(configPath)
	require.NoError(t, err)
	e := echo.New()
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://cinsear.ru", reader)
	q := req.URL.Query()
	q.Add("personsids", "3,4,5")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	qp := MapQueryParams(ctx)
	require.Equal(t, []int{3, 4, 5}, qp.PersonsIDs)
}
