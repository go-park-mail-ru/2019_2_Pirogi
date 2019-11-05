package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/makers"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
)

func (conn *MongoConnection) GetNextSequence(target string) (models.ID, error) {
	result := struct {
		Seq int `bson:"seq"`
	}{}
	err := conn.counters.FindOneAndUpdate(conn.context, bson.M{"_id": target},
		bson.M{"$inc": bson.M{"seq": 1}}).Decode(&result)
	return models.ID(result.Seq), errors.Wrap(err, "get next sequence failed")
}

func InsertUser(conn *MongoConnection, in models.NewUser) *models.Error {
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

func UpdateUser(conn *MongoConnection, in models.User) *models.Error {
	filter := bson.M{"usertrunc.id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.users.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "user not found")
	}
	return nil
}

// It is supposed that there cannot be films with the same title
func InsertFilm(conn *MongoConnection, in models.NewFilm) *models.Error {
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

func UpdateFilm(conn *MongoConnection, in models.Film) *models.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.films.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "film not found")
	}
	return nil
}

func UpsertUserCookie(conn *MongoConnection, in models.UserCookie) *models.Error {
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
func InsertPerson(conn *MongoConnection, in models.NewPerson) *models.Error {
	_, ok := conn.FindPersonByNameAndBirthday(in.Name, in.Birthday)
	if ok {
		return Error.New(http.StatusBadRequest, "person with this name and birthday already exists")
	}

	id, err := conn.GetNextSequence(configs.Default.PersonTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	newPerson, e := makers.MakePerson(id, in)
	if e != nil {
		return e
	}
	_, err = conn.persons.InsertOne(conn.context, newPerson)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	return nil
}

func UpdatePerson(conn *MongoConnection, in models.Person) *models.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.persons.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "person not found")
	}
	return nil
}

// TODO: check that user and film of review exist
func InsertReview(conn *MongoConnection, in models.NewReview) *models.Error {
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

func UpdateReview(conn *MongoConnection, in models.Review) *models.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.reviews.UpdateOne(conn.context, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "review not found")
	}
	return nil
}

func InsertStars(conn *MongoConnection, in models.Stars) *models.Error {
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

func InsertLike(conn *MongoConnection, in models.Like) *models.Error {
	_, err := conn.likes.InsertOne(conn.context, in)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert like in database")
	}
	return nil
}

func AggregateFilms(conn *MongoConnection, pipeline interface{}) ([]models.Film, *models.Error) {
	curs, err := conn.films.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, Error.New(http.StatusInternalServerError, "error while aggregating films")
	}
	var result []models.Film
	for curs.Next(conn.context) {
		f := models.Film{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, Error.New(http.StatusInternalServerError, "error while decoding aggregated result in films")
		}
		result = append(result, f)
	}
	return result, nil
}

func AggregateReviews(conn *MongoConnection, pipeline interface{}) ([]models.Review, *models.Error) {
	curs, err := conn.reviews.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, Error.New(http.StatusInternalServerError, "error while aggregating reviews")
	}
	var result []models.Review
	for curs.Next(conn.context) {
		f := models.Review{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, Error.New(http.StatusInternalServerError, "error while decoding aggregated result in reviews")
		}
		result = append(result, f)
	}
	return result, nil
}
