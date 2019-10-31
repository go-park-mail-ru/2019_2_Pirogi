package database

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

type Database interface {
	InsertOrUpdate(in interface{}) *models.Error
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

	FindPersonByNameAndBirthday(name string, birthday string) (models.Person, bool)
	FindPersonByID(id models.ID) (models.Person, bool)

	GetFilmsSortedByRating(limit int, offset int) ([]models.Film, *models.Error)
	GetFilmsOfGenreSortedByRating(genre models.Genre, limit int, offset int) ([]models.Film, *models.Error)
	GetFilmsOfYearSortedByRating(year int, limit int, offset int) ([]models.Film, *models.Error)

	GetReviewsSortedByDate(limit int, offset int) ([]models.Review, *models.Error)
	GetReviewsOfFilmSortedByDate(filmTitle string, limit int, offset int) ([]models.Review, *models.Error)
}
