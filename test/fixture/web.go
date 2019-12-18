package fixture

import (
	"bufio"
	"bytes"
	"github.com/labstack/echo"
	"net/http/httptest"
)

func NewEchoContext(body []byte, params map[string]string) echo.Context {
	e := echo.New()
	reader := bufio.NewReader(bytes.NewBuffer(body))
	req := httptest.NewRequest("POST", "http://cinsear.ru", reader)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	for k, v := range params {
		ctx.QueryParams().Add(k, v)
	}
	print(ctx.QueryParams())
	return ctx
}
