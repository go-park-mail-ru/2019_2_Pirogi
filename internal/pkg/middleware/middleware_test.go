package middleware

import (
	"bytes"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type DatabaseMock struct{}

func (DatabaseMock) Upsert(in interface{}) *models.Error {
	return nil
}

func (DatabaseMock) Get(id models.ID, target string) (interface{}, *models.Error) {
	panic("implement me")
}

func (DatabaseMock) Delete(in interface{}) *models.Error {
	return nil
}

func (DatabaseMock) ClearDB() {
	panic("implement me")
}

func (DatabaseMock) CheckCookie(cookie *http.Cookie) bool {
	panic("implement me")
}

func (DatabaseMock) FindUserByEmail(email string) (models.User, bool) {
	return models.User{
		Credentials: models.Credentials{
			Email:    "i@artbakulev.com",
			Password: "12345678",
		},
		UserTrunc: models.UserTrunc{},
	}, true
}

func (DatabaseMock) FindUserByID(id models.ID) (models.User, bool) {
	panic("implement me")
}

func (DatabaseMock) FindUserByCookie(cookie *http.Cookie) (models.User, bool) {
	if cookie.Value == "ok" {
		return models.User{}, true
	}
	return models.User{}, false
}

func (DatabaseMock) FindUsersByIDs(ids []models.ID) ([]models.User, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmByTitle(title string) (models.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmByID(id models.ID) (models.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmsByIDs(ids []models.ID) ([]models.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonByNameAndBirthday(name string, birthday string) (models.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonByID(id models.ID) (models.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonsByIDs(ids []models.ID) ([]models.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindReviewByID(id models.ID) (models.Review, bool) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsSortedByMark(limit int, offset int) ([]models.Film, *models.Error) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsOfGenreSortedByMark(genre models.Genre, limit int, offset int) ([]models.Film, *models.Error) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsOfYearSortedByMark(year string, limit int, offset int) ([]models.Film, *models.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsSortedByDate(limit int, offset int) ([]models.Review, *models.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsOfFilmSortedByDate(filmID models.ID, limit int, offset int) ([]models.Review, *models.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsOfAuthorSortedByDate(authorID models.ID, limit int, offset int) ([]models.Review, *models.Error) {
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
	h := AccessLogMiddleware(func(ctx echo.Context) error {
		return nil
	})
	e := echo.New()
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = h(c)
	require.NoError(t, err)

}
