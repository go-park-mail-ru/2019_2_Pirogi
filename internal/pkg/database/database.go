package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"net/http"
)

type Database interface {
	Upsert(in interface{}) *domains.Error
	Get(id domains.ID, target string) (interface{}, *domains.Error)
	Delete(in interface{}) *domains.Error
	ClearDB()

	CheckCookie(cookie *http.Cookie) bool

	FindUserByEmail(email string) (domains.User, bool)
	FindUserByID(id domains.ID) (domains.User, bool)
	FindUserByCookie(cookie *http.Cookie) (domains.User, bool)
	FindUsersByIDs(ids []domains.ID) ([]domains.User, bool)

	FindFilmByTitle(title string) (domains.Film, bool)
	FindFilmByID(id domains.ID) (domains.Film, bool)
	FindFilmsByIDs(ids []domains.ID) ([]domains.Film, bool)

	FindPersonByNameAndBirthday(name string, birthday string) (domains.Person, bool)
	FindPersonByID(id domains.ID) (domains.Person, bool)
	FindPersonsByIDs(ids []domains.ID) ([]domains.Person, bool)

	FindReviewByID(id domains.ID) (domains.Review, bool)

	GetByQuery(collectionName string, pipeline interface{}) ([]interface{}, *domains.Error)

	GetFilmsSortedByMark(limit, offset int) ([]domains.Film, *domains.Error)
	GetFilmsOfGenreSortedByMark(genre domains.Genre, limit, offset int) ([]domains.Film, *domains.Error)
	GetFilmsOfYearSortedByMark(year, limit, offset int) ([]domains.Film, *domains.Error)

	GetReviewsSortedByDate(limit, offset int) ([]domains.Review, *domains.Error)
	GetReviewsOfFilmSortedByDate(filmID domains.ID, limit, offset int) ([]domains.Review, *domains.Error)
	GetReviewsOfAuthorSortedByDate(authorID domains.ID, limit, offset int) ([]domains.Review, *domains.Error)
}
