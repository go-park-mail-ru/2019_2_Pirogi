package database

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

type Database interface {
	Upsert(in interface{}) *models.Error
	Get(id models.ID, target string) (interface{}, *models.Error)
	Delete(in interface{}) *models.Error
	ClearDB()

	CheckCookie(cookie *http.Cookie) bool

	FindUserByEmail(email string) (models.User, bool)
	FindUserByID(id models.ID) (models.User, bool)
	FindUserByCookie(cookie *http.Cookie) (models.User, bool)
	FindUsersByIDs(ids []models.ID) ([]models.User, bool)

	FindFilmByTitle(title string) (models.Film, bool)
	FindFilmByID(id models.ID) (models.Film, bool)
	FindFilmsByIDs(ids []models.ID) ([]models.Film, bool)

	FindPersonByNameAndBirthday(name string, birthday string) (models.Person, bool)
	FindPersonByID(id models.ID) (models.Person, bool)
	FindPersonsByIDs(ids []models.ID) ([]models.Person, bool)

	FindReviewByID(id models.ID) (models.Review, bool)

	FindListByID(id models.ID) (models.List, bool)

	GetByQuery(collectionName string, pipeline interface{}) ([]interface{}, *models.Error)

	GetFilmsSortedByMark(limit, offset int) ([]models.Film, *models.Error)
	GetFilmsOfGenreSortedByMark(genre models.Genre, limit, offset int) ([]models.Film, *models.Error)
	GetFilmsOfYearSortedByMark(year, limit, offset int) ([]models.Film, *models.Error)

	GetReviewsSortedByDate(limit, offset int) ([]models.Review, *models.Error)
	GetReviewsOfFilmSortedByDate(filmID models.ID, limit, offset int) ([]models.Review, *models.Error)
	GetReviewsOfAuthorSortedByDate(authorID models.ID, limit, offset int) ([]models.Review, *models.Error)
}
