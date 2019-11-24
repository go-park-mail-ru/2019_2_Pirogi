package database

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"net/http"
)

type MongoConnection struct {
	client  *mongo.Client
	context context.Context

	users         *mongo.Collection
	cookies       *mongo.Collection
	films         *mongo.Collection
	persons       *mongo.Collection
	likes         *mongo.Collection
	reviews       *mongo.Collection
	counters      *mongo.Collection
	subscriptions *mongo.Collection
}

func getMongoClient(mongoHost string) (*mongo.Client, error) {
	credentials := &options.Credential{
		Username:   configs.Default.MongoUser,
		Password:   configs.Default.MongoPwd,
		AuthSource: configs.Default.MongoDbName,
	}
	clientOpt := &options.ClientOptions{Auth: credentials}
	clientOpt.ApplyURI(mongoHost)
	client, err := mongo.NewClient(clientOpt)
	return client, err
}

func InitMongo(mongoHost string) (*MongoConnection, error) {
	client, err := getMongoClient(mongoHost)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	conn := MongoConnection{
		client:        client,
		context:       context.Background(),
		users:         client.Database(configs.Default.MongoDbName).Collection(configs.Default.UsersCollectionName),
		cookies:       client.Database(configs.Default.MongoDbName).Collection(configs.Default.CookiesCollectionName),
		films:         client.Database(configs.Default.MongoDbName).Collection(configs.Default.FilmsCollectionName),
		persons:       client.Database(configs.Default.MongoDbName).Collection(configs.Default.PersonsCollectionName),
		likes:         client.Database(configs.Default.MongoDbName).Collection(configs.Default.LikesCollectionName),
		reviews:       client.Database(configs.Default.MongoDbName).Collection(configs.Default.ReviewsCollectionName),
		counters:      client.Database(configs.Default.MongoDbName).Collection(configs.Default.CountersCollectionName),
		subscriptions: client.Database(configs.Default.MongoDbName).Collection(configs.Default.SubscriptionCollectionName),
	}

	return &conn, err
}

func (conn *MongoConnection) InitCounters() error {
	_, err := conn.counters.InsertMany(conn.context, []interface{}{
		bson.M{"_id": configs.Default.UserTargetName, "seq": 0},
		bson.M{"_id": configs.Default.FilmTargetName, "seq": 0},
		bson.M{"_id": configs.Default.PersonTargetName, "seq": 0},
		bson.M{"_id": configs.Default.ReviewTargetName, "seq": 0},
	})
	return errors.Wrap(err, "init counters collection failed")
}

func (conn *MongoConnection) Upsert(in interface{}) *model.Error {
	var e *model.Error
	switch in := in.(type) {
	case model.UserNew:
		e = InsertUser(conn, in)
	case model.User:
		e = UpdateUser(conn, in)
	case model.FilmNew:
		e = InsertFilm(conn, in)
	case model.Film:
		e = UpdateFilm(conn, in)
	case model.Cookie:
		e = UpsertUserCookie(conn, in)
	case model.PersonNew:
		e = InsertPerson(conn, in)
	case model.Person:
		e = UpdatePerson(conn, in)
	case model.ReviewNew:
		e = InsertReview(conn, in)
	case model.Review:
		e = UpdateReview(conn, in)
	case model.Stars:
		e = InsertStars(conn, in)
	case model.Like:
		e = InsertLike(conn, in)
	case model.SubscriptionNew:
		e = InsertSubscription(conn, in)
	case model.Subscription:
		e = UpdateSubscription(conn, in)
	default:
		e = model.NewError(http.StatusBadRequest, "Данный тип данных не поддерживается базой данных")
	}
	return e
}

func (conn *MongoConnection) Get(id model.ID, target string) (interface{}, *model.Error) {
	switch target {
	case configs.Default.UserTargetName:
		u, err := conn.FindUserByID(id)
		if err != nil {
			return nil, err
		}
		return u, nil
	case configs.Default.FilmTargetName:
		f, ok := conn.FindFilmByID(id)
		if ok {
			return f, nil
		}
		return nil, model.NewError(http.StatusNotFound, "no film with the id: "+id.String())
	case configs.Default.PersonTargetName:
		f, ok := conn.FindPersonByID(id)
		if ok {
			return f, nil
		}
		return nil, model.NewError(http.StatusNotFound, "no person with the id: "+id.String())
	case configs.Default.SubscriptionTargetName:
		f, err := conn.FindSubscriptionByUserID(id)
		if err == nil {
			return f, nil
		}
		return nil, model.NewError(http.StatusNotFound, "Подписки не существует. ID: "+id.String())
	}
	return nil, model.NewError(http.StatusNotFound, "Не поддерживаемые тип данных: "+target)
}

func (conn *MongoConnection) FindSubscriptionByUserID(userID model.ID) (model.Subscription, *model.Error) {
	result := model.Subscription{}
	err := conn.subscriptions.FindOne(conn.context, bson.M{"userid": userID}).Decode(&result)
	if err != nil {
		return model.Subscription{}, model.NewError(404, err.Error())
	}
	return result, nil
}

func (conn *MongoConnection) Delete(in interface{}) *model.Error {
	switch in := in.(type) {
	case http.Cookie:
		u, err := conn.FindUserByCookie(&in)
		if err != nil {
			return err
		}
		_, e := conn.cookies.DeleteOne(conn.context, bson.M{"_id": u.ID})
		if e != nil {
			return model.NewError(http.StatusInternalServerError, "cannot delete cookie from database: ", e.Error())
		}
		return nil
	}
	return model.NewError(http.StatusBadRequest, "not supported type")
}

func (conn *MongoConnection) ClearDB() {
	_, _ = conn.users.DeleteMany(conn.context, bson.M{})
	_, _ = conn.cookies.DeleteMany(conn.context, bson.M{})
	_, _ = conn.films.DeleteMany(conn.context, bson.M{})
	_, _ = conn.counters.DeleteMany(conn.context, bson.M{})
	_, _ = conn.reviews.DeleteMany(conn.context, bson.M{})
	_, _ = conn.persons.DeleteMany(conn.context, bson.M{})
	_, _ = conn.likes.DeleteMany(conn.context, bson.M{})
}

func (conn *MongoConnection) CheckCookie(cookie *http.Cookie) bool {
	foundCookie := model.Cookie{}
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	return err == nil
}

func (conn *MongoConnection) FindUserByEmail(email string) (model.User, *model.Error) {
	result := model.User{}
	err := conn.users.FindOne(conn.context, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return model.User{}, model.NewError(404, err.Error())
	}
	return result, nil
}

func (conn *MongoConnection) FindUserByID(id model.ID) (model.User, *model.Error) {
	result := model.User{}
	err := conn.users.FindOne(conn.context, bson.M{"id": id}).Decode(&result)
	if err != nil {
		return model.User{}, model.NewError(404, "Пользователя с таким ID не существует")
	}
	return result, nil
}

func (conn *MongoConnection) FindUserByCookie(cookie *http.Cookie) (model.User, *model.Error) {
	foundCookie := model.Cookie{}
	zap.S().Debug(cookie)
	zap.S().Debug(conn.context)
	zap.S().Debug(foundCookie)
	zap.S().Debug(cookie.Value)
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	zap.S().Debug(err)
	zap.S().Debug(foundCookie)
	if err != nil {
		return model.User{}, model.NewError(404, err.Error())
	}
	return conn.FindUserByID(foundCookie.UserID)
}

func (conn *MongoConnection) FindUsersByIDs(ids []model.ID) []model.User {
	var result []model.User
	for _, id := range ids {
		u, err := conn.FindUserByID(id)
		if err != nil {
			continue
		}
		result = append(result, u)
	}
	return result
}

//TODO: надо будет переделать на один запрос к базе)))
func (conn *MongoConnection) FindPersonsByIDs(ids []model.ID) ([]model.Person, bool) {
	var result []model.Person
	ok := true
	for _, id := range ids {
		u, found := conn.FindPersonByID(id)
		if !found {
			ok = found
			continue
		}
		result = append(result, u)
	}
	return result, ok
}

func (conn *MongoConnection) FindFilmsByIDs(ids []model.ID) []model.Film {
	var result []model.Film
	for _, id := range ids {
		u, found := conn.FindFilmByID(id)
		if !found {
			continue
		}
		result = append(result, u)
	}
	return result
}

func (conn *MongoConnection) FindFilmByTitle(title string) (model.Film, bool) {
	result := model.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"title": title}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindFilmByID(id model.ID) (model.Film, bool) {
	result := model.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindPersonByNameAndBirthday(name string, birthday string) (model.Person, bool) {
	result := model.Person{}
	err := conn.persons.FindOne(conn.context, bson.M{"name": name, "birthday": birthday}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindPersonByID(id model.ID) (model.Person, bool) {
	result := model.Person{}
	err := conn.persons.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindReviewByID(id model.ID) (model.Review, bool) {
	result := model.Review{}
	err := conn.reviews.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindSubscriptionsOnPerson(personID model.ID) ([]model.Subscription, *model.Error) {
	return AggregateSubscriptions(conn, personID)
}

func (conn *MongoConnection) GetByQuery(collectionName string, pipeline interface{}) ([]interface{}, *model.Error) {
	switch collectionName {
	case configs.Default.FilmsCollectionName:
		return AggregateFilms(conn, pipeline)
	case configs.Default.PersonsCollectionName:
		return AggregatePersons(conn, pipeline)
	default:
		return nil, model.NewError(http.StatusBadRequest, "not supported type")
	}
}

func (conn *MongoConnection) GetFilmsSortedByMark(limit, offset int) ([]model.Film, *model.Error) {
	pipeline := []bson.M{
		{"$sort": bson.M{"mark": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	films, err := AggregateFilms(conn, pipeline)
	if err != nil {
		return nil, err
	}
	return FromInterfaceToFilm(films), nil
}

func (conn *MongoConnection) GetFilmsOfGenreSortedByMark(genre model.Genre, limit, offset int) ([]model.Film, *model.Error) {
	pipeline := []bson.M{
		{"$match": bson.M{"genres": genre}},
		{"$sort": bson.M{"mark": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	films, err := AggregateFilms(conn, pipeline)
	if err != nil {
		return nil, err
	}
	return FromInterfaceToFilm(films), nil
}

func (conn *MongoConnection) GetFilmsOfYearSortedByMark(year, limit, offset int) ([]model.Film, *model.Error) {
	pipeline := []bson.M{
		{"$match": bson.M{"year": year}},
		{"$sort": bson.M{"mark": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	films, err := AggregateFilms(conn, pipeline)
	if err != nil {
		return nil, err
	}
	return FromInterfaceToFilm(films), nil
}

func (conn *MongoConnection) GetReviewsSortedByDate(limit, offset int) ([]model.Review, *model.Error) {
	pipeline := []bson.M{
		{"$sort": bson.M{"date": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	return AggregateReviews(conn, pipeline)
}

func (conn *MongoConnection) GetReviewsOfFilmSortedByDate(filmID model.ID, limit, offset int) ([]model.Review, *model.Error) {
	pipeline := []bson.M{
		{"$match": bson.M{"filmid": filmID}},
		{"$sort": bson.M{"date": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	return AggregateReviews(conn, pipeline)
}

func (conn *MongoConnection) GetReviewsOfAuthorSortedByDate(authorID model.ID, limit, offset int) ([]model.Review, *model.Error) {
	pipeline := []bson.M{
		{"$match": bson.M{"authorid": authorID}},
		{"$sort": bson.M{"date": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	return AggregateReviews(conn, pipeline)
}
