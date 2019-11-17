package auth

import (
	"bufio"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/models"
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
