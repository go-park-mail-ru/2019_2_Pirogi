package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const configsPath = "../../../configs"

func TestSetCSRFCookie(t *testing.T) {
	err := configuration.UnmarshalConfigs(configsPath)
	require.NoError(t, err)
	e := echo.New()
	buf := new(bytes.Buffer)
	e.Logger.SetOutput(buf)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := SetCSRFCookie(func(c echo.Context) error {
		return nil
	})
	err = h(c)
	require.NoError(t, err)
}

func TestHeaderMiddleware(t *testing.T) {
	err := configuration.UnmarshalConfigs(configsPath)
	require.NoError(t, err)
	e := echo.New()
	buf := new(bytes.Buffer)
	e.Logger.SetOutput(buf)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	var newRec http.ResponseWriter
	h := HeaderMiddleware(func(ctx echo.Context) error {
		newRec = ctx.Response()
		return nil
	})
	err = h(c)
	require.NoError(t, err)
	require.Equal(t, "*", newRec.Header().Get("Access-Control-Allow-Origin"))
}

func TestAccessLogMiddleware(t *testing.T) {
	err := configuration.UnmarshalConfigs(configsPath)
	require.NoError(t, err)
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	h := GetAccessLogMiddleware(logger)(func(ctx echo.Context) error {
		return nil
	})
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = h(c)
	require.NoError(t, err)

}
