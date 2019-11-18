package security

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestXSSFilterStrings(t *testing.T) {
	input := []string{"<script>alert('you have been pwned')</script>", "<script>console.log('he-he-he')</script>"}
	expected := []string{"&lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;", "&lt;script&gt;console.log(&#39;he-he-he&#39;)&lt;/script&gt;"}
	actual := XSSFilterStrings(input)
	require.Equal(t, expected, actual)
}

func TestXSSFilterGenres(t *testing.T) {
	input := []domains.Genre{"<script>alert('you have been pwned')</script>", "<script>console.log('he-he-he')</script>"}
	expected := []domains.Genre{"&lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;", "&lt;script&gt;console.log(&#39;he-he-he&#39;)&lt;/script&gt;"}
	actual := XSSFilterGenres(input)
	require.Equal(t, expected, actual)
}

func TestXSSFilterRoles(t *testing.T) {
	input := []domains.Role{"<script>alert('you have been pwned')</script>", "<script>console.log('he-he-he')</script>"}
	expected := []domains.Role{"&lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;", "&lt;script&gt;console.log(&#39;he-he-he&#39;)&lt;/script&gt;"}
	actual := XSSFilterRoles(input)
	require.Equal(t, expected, actual)
}

func TestCheckNoCSRFFail(t *testing.T) {
	e := echo.New()
	recorder := httptest.NewRecorder()
	cookie := &http.Cookie{
		Name:     configs.Default.CSRFCookieName,
		Value:    "test",
		Path:     "/",
		Expires:  time.Now().Add(configs.Default.CookieAuthDurationHours * time.Hour),
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(recorder, cookie)
	req := &http.Request{Header: http.Header{"Cookie": []string{recorder.Body.String()}}}
	req.Header.Set(configs.Default.CSRFHeader, "_csrf=invalid")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ok := CheckNoCSRF(c)
	require.False(t, ok)
}
