package usecase

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"

	"github.com/stretchr/testify/require"
)

const configsPath = "../../../configs"

type DatabaseMock struct{}

func (DatabaseMock) Upsert(in interface{}) *model.Error {
	return nil
}

func (DatabaseMock) Get(id model.ID, target string) (interface{}, *model.Error) {
	panic("implement me")
}

func (DatabaseMock) Delete(in interface{}) *model.Error {
	return nil
}

func (DatabaseMock) ClearDB() {
	panic("implement me")
}

func (DatabaseMock) CheckCookie(cookie *http.Cookie) bool {
	panic("implement me")
}

func (DatabaseMock) FindUserByEmail(email string) (model.User, bool) {
	return model.User{
		Email:    "i@artbakulev.com",
		Password: "12345678",
	}, true
}

func (DatabaseMock) FindUserByID(id model.ID) (model.User, bool) {
	panic("implement me")
}

func (DatabaseMock) FindUserByCookie(cookie *http.Cookie) (model.User, bool) {
	if cookie.Value == "ok" {
		return model.User{}, true
	}
	return model.User{}, false
}

func (DatabaseMock) FindUsersByIDs(ids []model.ID) ([]model.User, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmByTitle(title string) (model.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmByID(id model.ID) (model.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindFilmsByIDs(ids []model.ID) ([]model.Film, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonByNameAndBirthday(name string, birthday string) (model.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonByID(id model.ID) (model.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindPersonsByIDs(ids []model.ID) ([]model.Person, bool) {
	panic("implement me")
}

func (DatabaseMock) FindReviewByID(id model.ID) (model.Review, bool) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsSortedByMark(limit int, offset int) ([]model.Film, *model.Error) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsOfGenreSortedByMark(genre model.Genre, limit int, offset int) ([]model.Film, *model.Error) {
	panic("implement me")
}

func (DatabaseMock) GetFilmsOfYearSortedByMark(year string, limit int, offset int) ([]model.Film, *model.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsSortedByDate(limit int, offset int) ([]model.Review, *model.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsOfFilmSortedByDate(filmID model.ID, limit int, offset int) ([]model.Review, *model.Error) {
	panic("implement me")
}

func (DatabaseMock) GetReviewsOfAuthorSortedByDate(authorID model.ID, limit int, offset int) ([]model.Review, *model.Error) {
	panic("implement me")
}
func TestGenerateCookie(t *testing.T) {
	const cookieName = "test"
	const value = "testValue"
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(cookieName, value)
	expectedCookie := http.Cookie{
		Name:     cookieName,
		Value:    hash.SHA1(value),
		Expires:  cookie.Cookie.Expires,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	require.Equal(t, cookie.Cookie, expectedCookie)
}

func TestExpireCookie(t *testing.T) {
	err := configuration.UnmarshalConfigs(configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	const cookieName = "test"
	const value = "testValue"
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(cookieName, value)
	require.True(t, cookie.Cookie.Expires.After(time.Now()))
	cookie.Expire()
	require.True(t, cookie.Cookie.Expires.Before(time.Now()))
}

//func TestGetUserByRequestFail(t *testing.T) {
//	req := http.Request{}
//	_, err := GetUserByRequest(&req, DatabaseMock{})
//	require.NotNil(t, err)
//}
//
//func TestLogin(t *testing.T) {
//	e := echo.New()
//	req := http.Request{}
//	rec := httptest.NewRecorder()
//	ctx := e.NewContext(&req, rec)
//	err := Login(ctx, DatabaseMock{}, "i@artbakulev", "12345678")
//	require.Nil(t, err)
//}
//
//func TestLoginFail(t *testing.T) {
//	e := echo.New()
//	req := http.Request{}
//	rec := httptest.NewRecorder()
//	ctx := e.NewContext(&req, rec)
//	err := Login(ctx, DatabaseMock{}, "i@artbakulev", "invalid")
//	require.NotNil(t, err)
//}
//
//func TestLoginCheck(t *testing.T) {
//	e := echo.New()
//	reader := bufio.NewReader(nil)
//	req := httptest.NewRequest("POST", "http://www.google.com", reader)
//	cookie := &http.Cookie{
//		Name:  "cinsear_session",
//		Value: "ok",
//		Path:  "/",
//	}
//	req.AddCookie(cookie)
//	rec := httptest.NewRecorder()
//	ctx := e.NewContext(req, rec)
//	ok := LoginCheck(ctx, DatabaseMock{})
//	require.True(t, ok)
//}
//
//func TestGetUserByRequest(t *testing.T) {
//	reader := bufio.NewReader(nil)
//	req := httptest.NewRequest("POST", "http://www.google.com", reader)
//	cookie := &http.Cookie{
//		Name:  "cinsear_session",
//		Value: "ok",
//		Path:  "/",
//	}
//	req.AddCookie(cookie)
//	_, ok := GetUserByRequest(req, DatabaseMock{})
//	require.True(t, ok)
//}
//
//func TestLogout(t *testing.T) {
//	e := echo.New()
//	reader := bufio.NewReader(nil)
//	req := httptest.NewRequest("POST", "http://www.google.com", reader)
//	cookie := &http.Cookie{
//		Name:  "cinsear_session",
//		Value: "ok",
//		Path:  "/",
//	}
//	req.AddCookie(cookie)
//	rec := httptest.NewRecorder()
//	ctx := e.NewContext(req, rec)
//	err := Logout(ctx, DatabaseMock{})
//	require.Nil(t, err)
//}
