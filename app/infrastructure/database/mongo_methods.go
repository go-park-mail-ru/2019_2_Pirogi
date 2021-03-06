package database

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.mongodb.org/mongo-driver/bson"
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
	id, e := conn.GetNextSequence(configs.Default.UserTargetName)
	if e != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert user in database")
	}
	u := in.ToUser(id)
	_, e = conn.users.InsertOne(conn.context, u)
	if e != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert user in database")
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
	f := in.ToFilm(id)
	_, err = conn.films.InsertOne(conn.context, f)
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
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	newPerson := in.ToPerson(id)
	_, err = conn.persons.InsertOne(conn.context, newPerson)
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
	//foundUser, e := conn.Get(in.AuthorID, configs.Default.UserTargetName)
	//if e != nil {
	//	return model.NewError(http.StatusNotFound, "user not found")
	//}
	id, err := conn.GetNextSequence(configs.Default.ReviewTargetName)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert review in database")
	}
	rev := in.ToReview(id)
	_, err = conn.reviews.InsertOne(conn.context, rev)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert review in database")
	}
	//TODO: добавить поле, код внизу должен работать
	//u := foundUser.(model.User)
	//u.Reviews++
	//conn.Upsert(u)
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

func InsertRating(conn *MongoConnection, in model.Rating) *model.Error {
	_, err := conn.ratings.InsertOne(conn.context, in)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert rating in database")
	}
	return nil
}

func UpdateRating(conn *MongoConnection, in model.RatingUpdate) *model.Error {
	filter := bson.M{"userid": in.UserID, "filmid": in.FilmID}
	update := bson.M{"$set": bson.M{"mark": in.Stars}}
	_, err := conn.ratings.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "rating not found")
	}
	return nil
}

func InsertList(conn *MongoConnection, in model.ListNew) *model.Error {
	id, err := conn.GetNextSequence(configs.Default.ListTargetName)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert list in database")
	}
	list := in.ToList(id)
	_, e := conn.lists.InsertOne(conn.context, list)
	if e != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert list in database")
	}
	return nil
}

func UpdateList(conn *MongoConnection, in model.List) *model.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	_, err := conn.lists.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "list not found")
	}
	return nil
}

func InsertSubscription(conn *MongoConnection, in model.SubscriptionNew) *model.Error {
	_, err := conn.subscriptions.InsertOne(conn.context, in)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "cannot insert subscriptionNew in database")
	}
	return nil
}

func UpdateSubscription(conn *MongoConnection, in model.Subscription) *model.Error {
	filter := bson.M{"_id": in.UserID}
	update := bson.M{"$set": in}
	_, err := conn.subscriptions.UpdateOne(conn.context, filter, update)
	if err != nil {
		return model.NewError(http.StatusNotFound, "subscription not found")
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

func AggregateSubscriptions(conn *MongoConnection, personID model.ID) ([]model.Subscription, *model.Error) {
	curs, err := conn.subscriptions.Find(conn.context, bson.M{"personsid": bson.M{"$all": []model.ID{personID}}})
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "error while aggregating subscriptions", err.Error())
	}
	var result []model.Subscription
	for curs.Next(conn.context) {
		p := model.Subscription{}
		err = curs.Decode(&p)
		if err != nil {
			return nil, model.NewError(http.StatusInternalServerError, "error while decoding aggregated result in subscrtiptions", err.Error())
		}
		result = append(result, p)
	}
	zap.S().Debug(result)
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
