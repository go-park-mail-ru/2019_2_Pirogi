package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/review"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/user"
	"net/http"
)

type Database interface {
	Upsert(in interface{}) (model.ID, *model.Error)
	Get(id domains.ID, target string) (interface{}, *domains.Error)
	Delete(in interface{}) *domains.Error
	ClearDB()

	CheckCookie(cookie *http.Cookie) bool

	FindUserByEmail(email string) (user.User, bool)
	FindUserByID(id domains.ID) (user.User, bool)
	FindUserByCookie(cookie *http.Cookie) (user.User, bool)
	FindUsersByIDs(ids []domains.ID) ([]user.User, bool)

	FindFilmByTitle(title string) (film.Film, bool)
	FindFilmByID(id domains.ID) (film.Film, bool)
	FindFilmsByIDs(ids []model.ID) []model.Film

	FindPersonByNameAndBirthday(name string, birthday string) (domains.Person, bool)
	FindPersonByID(id domains.ID) (domains.Person, bool)
	FindPersonsByIDs(ids []domains.ID) ([]domains.Person, bool)

	FindReviewByID(id domains.ID) (review.Review, bool)

	GetByQuery(collectionName string, pipeline interface{}) ([]interface{}, *domains.Error)

	GetFilmsSortedByMark(limit, offset int) ([]film.Film, *domains.Error)
	GetFilmsOfGenreSortedByMark(genre domains.Genre, limit, offset int) ([]film.Film, *domains.Error)
	GetFilmsOfYearSortedByMark(year, limit, offset int) ([]film.Film, *domains.Error)

	GetReviewsSortedByDate(limit, offset int) ([]review.Review, *domains.Error)
	GetReviewsOfFilmSortedByDate(filmID domains.ID, limit, offset int) ([]review.Review, *domains.Error)
	GetReviewsOfAuthorSortedByDate(authorID domains.ID, limit, offset int) ([]review.Review, *domains.Error)
}
