package server

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
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

func TestCreateAPIServer(t *testing.T) {
	// here you can see the best test ever
	_, err := CreateAPIServer(DatabaseMock{})
	require.NoError(t, err)
}
