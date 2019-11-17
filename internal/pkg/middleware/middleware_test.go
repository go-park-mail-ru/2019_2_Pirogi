package middleware

import (
	"bytes"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type DatabaseMock struct{}

func (DatabaseMock) Upsert(in interface{}) *domains.Error {
	return nil
}

func (DatabaseMock) Get(id domains.ID, target string) (interface{}, *domains.Error) {
	panic("implement me")
}

func (DatabaseMock) Delete(in interface{}) *domains.Error {
	return nil
}

func (DatabaseMock) ClearDB() {
	panic("implement me")
}

func (DatabaseMock) CheckCookie(cookie *http.Cookie) bool {
	panic("implement me")
}

func (DatabaseMock) FindUserByEmail(email string) (domains.User, bool) {
	return domains.User{
		Credentials: models.Credentials{
			Email:    "i@artbakulev.com",
			Password: "12345678",
		},
		UserTrunc: domains.UserTrunc{},
	}, true
}

func (DatabaseMock) FindUserByID(id domains.ID) (domains.User, bool) {
	panic("implement me")
}

func (DatabaseMock) FindUserByCookie(cookie *http.Cookie) (domains.User, bool) {
	if cookie.Value == "ok" {
		return domains.User{}, true
	}
	return domains.User{}, false
}

func (DatabaseMock) FindUsersByIDs(ids []domains.ID) ([]domains.User, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmByTitle(title string) (domains.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmByID(id domains.ID) (domains.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmsByIDs(ids []domains.ID) ([]domains.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonByNameAndBirthday(name string, birthday string) (domains.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonByID(id domains.ID) (domains.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonsByIDs(ids []domains.ID) ([]domains.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindReviewByID(id domains.ID) (domains.Review, bool) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsSortedByMark(limit int, offset int) ([]domains.Film, *domains.Error) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsOfGenreSortedByMark(genre domains.Genre, limit int, offset int) ([]domains.Film, *domains.Error) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsOfYearSortedByMark(year string, limit int, offset int) ([]domains.Film, *domains.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsSortedByDate(limit int, offset int) ([]domains.Review, *domains.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsOfFilmSortedByDate(filmID domains.ID, limit int, offset int) ([]domains.Review, *domains.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsOfAuthorSortedByDate(authorID domains.ID, limit int, offset int) ([]domains.Review, *domains.Error) {
	panic("implement me")
}

func TestExpireInvalidCookiesMiddleware(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	e.Logger.SetOutput(buf)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	expiredCookie := &http.Cookie{
		Name:    "cinsear_session",
		Value:   "test",
		Path:    "",
		Domain:  "",
		Expires: time.Now(),
	}
	req.AddCookie(expiredCookie)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := ExpireInvalidCookiesMiddleware(DatabaseMock{})(func(c echo.Context) error {
		return nil
	})
	err := h(c)
	require.NoError(t, err)
	session, err := req.Cookie("cinsear_session")
	require.NoError(t, err)
	require.True(t, session.Expires.Before(time.Now()))
}

func TestSetCSRFCookie(t *testing.T) {
	configPath := "../../../configs"
	err := common.UnmarshalConfigs(&configPath)
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
	configPath := "../../../configs"
	err := common.UnmarshalConfigs(&configPath)
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
	configPath := "../../../configs"
	err := common.UnmarshalConfigs(&configPath)
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
