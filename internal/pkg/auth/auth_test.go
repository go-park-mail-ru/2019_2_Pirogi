package auth

import (
	"bufio"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/stretchr/testify/require"
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
func TestGenerateCookie(t *testing.T) {
	const cookieName = "test"
	const value = "testValue"

	generatedCookie := GenerateCookie(cookieName, value)
	expectedCookie := http.Cookie{
		Name:     cookieName,
		Value:    user.GetMD5Hash(value),
		Expires:  generatedCookie.Expires,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	require.Equal(t, generatedCookie, expectedCookie)
}

func TestExpireCookie(t *testing.T) {
	configsPath := "../../../configs"
	err := common.UnmarshalConfigs(&configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	const cookieName = "test"
	const value = "testValue"

	generatedCookie := GenerateCookie(cookieName, value)
	require.True(t, generatedCookie.Expires.After(time.Now()))
	ExpireCookie(&generatedCookie)
	require.True(t, generatedCookie.Expires.Before(time.Now()))
}

func TestGetUserByRequestFail(t *testing.T) {
	req := http.Request{}
	_, err := GetUserByRequest(&req, DatabaseMock{})
	require.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	e := echo.New()
	req := http.Request{}
	rec := httptest.NewRecorder()
	ctx := e.NewContext(&req, rec)
	err := Login(ctx, DatabaseMock{}, "i@artbakulev", "12345678")
	require.Nil(t, err)
}

func TestLoginFail(t *testing.T) {
	e := echo.New()
	req := http.Request{}
	rec := httptest.NewRecorder()
	ctx := e.NewContext(&req, rec)
	err := Login(ctx, DatabaseMock{}, "i@artbakulev", "invalid")
	require.NotNil(t, err)
}

func TestLoginCheck(t *testing.T) {
	e := echo.New()
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://www.google.com", reader)
	cookie := &http.Cookie{
		Name:  "cinsear_session",
		Value: "ok",
		Path:  "/",
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ok := LoginCheck(ctx, DatabaseMock{})
	require.True(t, ok)
}

func TestGetUserByRequest(t *testing.T) {
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://www.google.com", reader)
	cookie := &http.Cookie{
		Name:  "cinsear_session",
		Value: "ok",
		Path:  "/",
	}
	req.AddCookie(cookie)
	_, ok := GetUserByRequest(req, DatabaseMock{})
	require.True(t, ok)
}

func TestLogout(t *testing.T) {
	e := echo.New()
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://www.google.com", reader)
	cookie := &http.Cookie{
		Name:  "cinsear_session",
		Value: "ok",
		Path:  "/",
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	err := Logout(ctx, DatabaseMock{})
	require.Nil(t, err)
}
