package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/like"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/review"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/stars"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/models"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
)

func (conn *MongoConnection) GetNextSequence(target string) (domains.ID, error) {
	result := struct {
		Seq int `bson:"seq"`
	}{}
	err := conn.counters.FindOneAndUpdate(conn.context, bson.M{"_id": target},
		bson.M{"$inc": bson.M{"seq": 1}}).Decode(&result)
	return domains.ID(result.Seq), errors.Wrap(err, "get next sequence failed")
}

func InsertUser(conn *MongoConnection, in domains.UserNew) *domains.Error {
	_, ok := conn.FindUserByEmail(in.Email)
	if ok {
		return Error.New(http.StatusBadRequest, "user with the email already exists")
	}
	id, err := conn.GetNextSequence(configs.Default.UserTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert user in database")
	}
	u, e := user.CreateUser(id, &in)
	if e != nil {
		return e
	}
	_, err = conn.users.InsertOne(conn.context, u)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert user in database")
	}
	return nil
}

func UpdateUser(conn *MongoConnection, in domains.User) *domains.Error {
	filter := bson.M{"usertrunc.id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.users.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "user not found")
	}
	return nil
}

// It is supposed that there cannot be films with the same title
func InsertFilm(conn *MongoConnection, in film.FilmNew) *domains.Error {
	_, ok := conn.FindFilmByTitle(in.Title)
	if ok {
		return Error.New(http.StatusBadRequest, "film with the title already exists")
	}
	id, err := conn.GetNextSequence(configs.Default.FilmTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert film in database")
	}
	f := makers.MakeFilm(id, &in)
	_, err = conn.films.InsertOne(conn.context, f)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert film in database")
	}
	return nil
}

func UpdateFilm(conn *MongoConnection, in film.Film) *domains.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.films.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "film not found")
	}
	return nil
}

func UpsertUserCookie(conn *MongoConnection, in models.UserCookie) *domains.Error {
	filter := bson.M{"_id": in.UserID}
	foundCookie := models.UserCookie{}
	err := conn.cookies.FindOne(conn.context, filter).Decode(&foundCookie)
	if err != nil {
		_, err = conn.cookies.InsertOne(conn.context, in)
	} else {
		update := bson.M{"$set": in}
		_, err = conn.cookies.UpdateOne(conn.context, filter, update)
	}
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert cookie in database")
	}
	return nil
}

// It is supposed that there cannot be persons with the same name and birthday
func InsertPerson(conn *MongoConnection, in domains.PersonNew) *domains.Error {
	_, ok := conn.FindPersonByNameAndBirthday(in.Name, in.Birthday)
	if ok {
		return Error.New(http.StatusBadRequest, "person with this name and birthday already exists")
	}

	id, err := conn.GetNextSequence(configs.Default.PersonTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	newPerson := makers.MakePerson(id, in)
	_, err = conn.persons.InsertOne(conn.context, newPerson)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	return nil
}

func UpdatePerson(conn *MongoConnection, in domains.Person) *domains.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.persons.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "person not found")
	}
	return nil
}

// TODO: check that user and film of review exist
func InsertReview(conn *MongoConnection, in review.ReviewNew) *domains.Error {
	id, err := conn.GetNextSequence(configs.Default.ReviewTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert review in database")
	}
	rev := makers.MakeReview(id, in)
	_, err = conn.reviews.InsertOne(conn.context, rev)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert review in database")
	}
	return nil
}

func UpdateReview(conn *MongoConnection, in review.Review) *domains.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.reviews.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "review not found")
	}
	return nil
}

func InsertStars(conn *MongoConnection, in stars.Stars) *domains.Error {
	filter := bson.M{"_id": in.FilmID}
	// TODO: рассчитывать newMark
	newMark := in.Mark
	update := bson.M{"$set": bson.M{"mark": newMark}}
	_, err := conn.films.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "film not found")
	}
	return nil
}

func InsertLike(conn *MongoConnection, in like.Like) *domains.Error {
	_, err := conn.likes.InsertOne(conn.context, in)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert like in database")
	}
	return nil
}

func AggregateFilms(conn *MongoConnection, pipeline interface{}) ([]interface{}, *domains.Error) {
	curs, err := conn.films.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, Error.New(http.StatusInternalServerError, "error while aggregating films", err.Error())
	}
	var result []interface{}
	for curs.Next(conn.context) {
		f := film.Film{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, Error.New(http.StatusInternalServerError, "error while decoding aggregated result in films", err.Error())
		}
		result = append(result, f)
	}
	return result, nil
}

func AggregatePersons(conn *MongoConnection, pipeline interface{}) ([]interface{}, *domains.Error) {
	curs, err := conn.persons.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, Error.New(http.StatusInternalServerError, "error while aggregating persons", err.Error())
	}
	var result []interface{}
	for curs.Next(conn.context) {
		p := domains.Person{}
		err = curs.Decode(&p)
		if err != nil {
			return nil, Error.New(http.StatusInternalServerError, "error while decoding aggregated result in persons", err.Error())
		}
		result = append(result, p)
	}
	return result, nil
}

func AggregateReviews(conn *MongoConnection, pipeline interface{}) ([]review.Review, *domains.Error) {
	curs, err := conn.reviews.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, Error.New(http.StatusInternalServerError, "error while aggregating reviews")
	}
	var result []review.Review
	for curs.Next(conn.context) {
		f := review.Review{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, Error.New(http.StatusInternalServerError, "error while decoding aggregated result in reviews")
		}
		result = append(result, f)
	}
	return result, nil
}

func FromInterfaceToFilm(films []interface{}) []film.Film {
	result := make([]film.Film, len(films))
	for i, f := range films {
		result[i] = f.(film.Film)
	}
	return result
}
