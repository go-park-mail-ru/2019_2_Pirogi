package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"net/http"
)

type Database interface {
	Upsert(in interface{}) *model.Error
	Get(id model.ID, target string) (interface{}, *model.Error)
	Delete(in interface{}) *model.Error
	ClearDB()

	CheckCookie(cookie *http.Cookie) bool

	FindUserByEmail(email string) (model.User, *model.Error)
	FindUserByID(id model.ID) (model.User, *model.Error)
	FindUserByCookie(cookie *http.Cookie) (model.User, *model.Error)
	FindUsersByIDs(ids []model.ID) []model.User

	FindFilmsByIDs(ids []model.ID) []model.Film
	FindFilmByTitle(title string) (model.Film, bool)
	FindFilmByID(id model.ID) (model.Film, bool)

	FindPersonByNameAndBirthday(name string, birthday string) (model.Person, bool)
	FindPersonByID(id model.ID) (model.Person, bool)
	FindPersonsByIDs(ids []model.ID) ([]model.Person, bool)

	FindReviewByID(id model.ID) (model.Review, bool)

	FindListByID(id model.ID) (model.List, bool)
	FindListsByUserID(userId model.ID) ([]model.List, *model.Error)

	GetByQuery(collectionName string, pipeline interface{}) ([]interface{}, *model.Error)

	GetFilmsSortedByMark(limit, offset int) ([]model.Film, *model.Error)
	GetFilmsOfGenreSortedByMark(genre model.Genre, limit, offset int) ([]model.Film, *model.Error)
	GetFilmsOfYearSortedByMark(year, limit, offset int) ([]model.Film, *model.Error)

	GetReviewsSortedByDate(limit, offset int) ([]model.Review, *model.Error)
	GetReviewsOfFilmSortedByDate(filmID model.ID, limit, offset int) ([]model.Review, *model.Error)
	GetReviewsOfAuthorSortedByDate(authorID model.ID, limit, offset int) ([]model.Review, *model.Error)

	FindSubscriptionByUserID(userID model.ID) (model.Subscription, *model.Error)
	FindSubscriptionsOnPerson(personID model.ID) ([]model.Subscription, *model.Error)
}
