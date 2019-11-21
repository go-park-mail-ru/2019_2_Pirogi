package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"net/http"
)

func (conn *MongoConnection) GetNextSequence(target string) (model.ID, error) {
	result := struct {
		Seq int `bson:"seq"`
	}{}
	err := conn.counters.FindOneAndUpdate(conn.context, bson.M{"_id": target},
		bson.M{"$inc": bson.M{"seq": 1}}).Decode(&result)
	return model.ID(result.Seq), errors.Wrap(err, "get next sequence failed")
}

func InsertUser(conn *MongoConnection, in model.UserNew) *model.Error {
	_, err := conn.FindUserByEmail(in.Email)
	if err == nil {
		return model.NewError(http.StatusBadRequest, "user with the email is already existed")
	}
	id, e := conn.GetNextSequence(configs.Default.UserTargetName)
	if e != nil {
		return model.NewError(http.StatusInternalServerError, e.Error())
	}
	user := in.ToUser(id)
	_, e = conn.users.InsertOne(conn.context, user)
	if e != nil {
		return model.NewError(http.StatusInternalServerError, e.Error())
	}
	return nil
}

func UpdateUser(conn *MongoConnection, in model.User) *model.Error {
	filter := bson.M{"id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.users.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "user not found")
	}
	return nil
}

// It is supposed that there cannot be films with the same title
func InsertFilm(conn *MongoConnection, in model.FilmNew) *model.Error {
	_, ok := conn.FindFilmByTitle(in.Title)
	if ok {
		return model.NewError(http.StatusBadRequest, "film with the title already exists")
	}
	id, err := conn.GetNextSequence(configs.Default.FilmTargetName)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert film in database")
	}
	film := in.ToFilm(id)
	_, err = conn.films.InsertOne(conn.context, film)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert film in database")
	}
	return nil
}

func UpdateFilm(conn *MongoConnection, in model.Film) *model.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.films.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "film not found")
	}
	return nil
}

func UpsertUserCookie(conn *MongoConnection, in model.Cookie) *model.Error {
	filter := bson.M{"_id": in.UserID}
	foundCookie := model.Cookie{}
	err := conn.cookies.FindOne(conn.context, filter).Decode(&foundCookie)
	if err != nil {
		_, err = conn.cookies.InsertOne(conn.context, in)
	} else {
		update := bson.M{"$set": in}
		_, err = conn.cookies.UpdateOne(conn.context, filter, update)
	}
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert cookie in database")
	}
	return nil
}

// It is supposed that there cannot be persons with the same name and birthday
func InsertPerson(conn *MongoConnection, in model.PersonNew) *model.Error {
	_, ok := conn.FindPersonByNameAndBirthday(in.Name, in.Birthday)
	if ok {
		return model.NewError(http.StatusBadRequest, "person with this name and birthday already exists")
	}

	id, err := conn.GetNextSequence(configs.Default.PersonTargetName)
	zap.S().Debug(id)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	person := in.ToPerson(id)
	_, err = conn.persons.InsertOne(conn.context, person)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	return nil
}

func UpdatePerson(conn *MongoConnection, in model.Person) *model.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.persons.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "person not found")
	}
	return nil
}

// TODO: check that user and film of review exist
func InsertReview(conn *MongoConnection, in model.ReviewNew) *model.Error {
	id, err := conn.GetNextSequence(configs.Default.ReviewTargetName)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert review in database")
	}
	rev := in.ToReview(id)
	_, err = conn.reviews.InsertOne(conn.context, rev)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert review in database")
	}
	return nil
}

func UpdateReview(conn *MongoConnection, in model.Review) *model.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.reviews.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "review not found")
	}
	return nil
}

func InsertStars(conn *MongoConnection, in model.Stars) *model.Error {
	filter := bson.M{"_id": in.FilmID}
	// TODO: рассчитывать newMark
	newMark := in.Mark
	update := bson.M{"$set": bson.M{"mark": newMark}}
	_, err := conn.films.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "film not found")
	}
	return nil
}

func InsertLike(conn *MongoConnection, in model.Like) *model.Error {
	_, err := conn.likes.InsertOne(conn.context, in)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert like in database")
	}
	return nil
}

func AggregateFilms(conn *MongoConnection, pipeline interface{}) ([]interface{}, *model.Error) {
	curs, err := conn.films.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "error while aggregating films", err.Error())
	}
	var result []interface{}
	for curs.Next(conn.context) {
		f := model.Film{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, model.NewError(http.StatusInternalServerError, "error while decoding aggregated result in films", err.Error())
		}
		result = append(result, f)
	}
	return result, nil
}

func AggregatePersons(conn *MongoConnection, pipeline interface{}) ([]interface{}, *model.Error) {
	curs, err := conn.persons.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "error while aggregating persons", err.Error())
	}
	var result []interface{}
	for curs.Next(conn.context) {
		p := model.Person{}
		err = curs.Decode(&p)
		if err != nil {
			return nil, model.NewError(http.StatusInternalServerError, "error while decoding aggregated result in persons", err.Error())
		}
		result = append(result, p)
	}
	return result, nil
}

func AggregateReviews(conn *MongoConnection, pipeline interface{}) ([]model.Review, *model.Error) {
	curs, err := conn.reviews.Aggregate(conn.context, pipeline)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "error while aggregating reviews")
	}
	var result []model.Review
	for curs.Next(conn.context) {
		f := model.Review{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, model.NewError(http.StatusInternalServerError, "error while decoding aggregated result in reviews")
		}
		result = append(result, f)
	}
	return result, nil
}

func FromInterfaceToFilm(films []interface{}) []model.Film {
	result := make([]model.Film, len(films))
	for i, f := range films {
		result[i] = f.(model.Film)
	}
	return result
}
