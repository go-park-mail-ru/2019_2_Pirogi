package database

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os/user"
)

type MongoConnection struct {
	client  *mongo.Client
	context context.Context

	users    *mongo.Collection
	cookies  *mongo.Collection
	films    *mongo.Collection
	persons  *mongo.Collection
	likes    *mongo.Collection
	reviews  *mongo.Collection
	counters *mongo.Collection
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
		client:   client,
		context:  context.Background(),
		users:    client.Database(configs.Default.MongoDbName).Collection(configs.Default.UsersCollectionName),
		cookies:  client.Database(configs.Default.MongoDbName).Collection(configs.Default.CookiesCollectionName),
		films:    client.Database(configs.Default.MongoDbName).Collection(configs.Default.FilmsCollectionName),
		persons:  client.Database(configs.Default.MongoDbName).Collection(configs.Default.PersonsCollectionName),
		likes:    client.Database(configs.Default.MongoDbName).Collection(configs.Default.LikesCollectionName),
		reviews:  client.Database(configs.Default.MongoDbName).Collection(configs.Default.ReviewsCollectionName),
		counters: client.Database(configs.Default.MongoDbName).Collection(configs.Default.CountersCollectionName),
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

func (conn *MongoConnection) Upsert(in interface{}) *domains.Error {
	var e *domains.Error
	switch in := in.(type) {
	case domains.UserNew:
		e = InsertUser(conn, in)
	case domains.User:
		e = UpdateUser(conn, in)
	case domains.FilmNew:
		e = InsertFilm(conn, in)
	case domains.Film:
		e = UpdateFilm(conn, in)
	case domains.Cookie:
		e = UpsertUserCookie(conn, in)
	case domains.PersonNew:
		e = InsertPerson(conn, in)
	case domains.Person:
		e = UpdatePerson(conn, in)
	case domains.ReviewNew:
		e = InsertReview(conn, in)
	case domains.Review:
		e = UpdateReview(conn, in)
	case domains.Stars:
		e = InsertStars(conn, in)
	case domains.Like:
		e = InsertLike(conn, in)
	default:
		e = Error.NewError(http.StatusBadRequest, "not supported type")
	}
	return e
}

func (conn *MongoConnection) Get(id domains.ID, target string) (interface{}, *domains.Error) {
	switch target {
	case configs.Default.UserTargetName:
		u, ok := conn.FindUserByID(id)
		if ok {
			return u, nil
		}
		return nil, Error.NewError(http.StatusNotFound, "no user with id: "+id.String())
	case configs.Default.FilmTargetName:
		f, ok := conn.FindFilmByID(id)
		if ok {
			return f, nil
		}
		return nil, Error.NewError(http.StatusNotFound, "no film with the id: "+id.String())
	case configs.Default.PersonTargetName:
		f, ok := conn.FindPersonByID(id)
		if ok {
			return f, nil
		}
		return nil, Error.NewError(http.StatusNotFound, "no person with the id: "+id.String())
	}
	return nil, Error.NewError(http.StatusNotFound, "not supported type: "+target)
}

func (conn *MongoConnection) Delete(in interface{}) *domains.Error {
	switch in := in.(type) {
	case http.Cookie:
		u, ok := conn.FindUserByCookie(&in)
		if !ok {
			return nil
		}
		_, err := conn.cookies.DeleteOne(conn.context, bson.M{"_id": })
		if err != nil {
			return Error.NewError(http.StatusInternalServerError, "cannot delete cookie from database: "+err.Error())
		}
		return nil
	}
	return Error.NewError(http.StatusBadRequest, "not supported type")
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
	foundCookie := models.UserCookie{}
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	return err == nil
}

func (conn *MongoConnection) FindUserByEmail(email string) (user.User, bool) {
	result := user.User{}
	err := conn.users.FindOne(conn.context, bson.M{"credentials.email": email}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindUserByID(id domains.ID) (user.User, bool) {
	result := user.User{}
	err := conn.users.FindOne(conn.context, bson.M{"usertrunc.id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindUserByCookie(cookie *http.Cookie) (domains.User, bool) {
	foundCookie := cookie.
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	if err != nil {
		return user.User{}, false
	}
	return conn.FindUserByID(foundCookie.UserID)
}

func (conn *MongoConnection) FindUsersByIDs(ids []domains.ID) ([]user.User, bool) {
	var result []user.User
	ok := true
	for _, id := range ids {
		u, found := conn.FindUserByID(id)
		if !found {
			ok = found
			continue
		}
		result = append(result, u)
	}
	return result, ok
}

//TODO: надо будет переделать на один запрос к базе)))
func (conn *MongoConnection) FindPersonsByIDs(ids []domains.ID) ([]domains.Person, bool) {
	var result []domains.Person
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

func (conn *MongoConnection) FindFilmsByIDs(ids []domains.ID) ([]film.Film, bool) {
	var result []film.Film
	ok := true
	for _, id := range ids {
		u, found := conn.FindFilmByID(id)
		if !found {
			ok = found
			continue
		}
		result = append(result, u)
	}
	return result, ok
}

func (conn *MongoConnection) FindFilmByTitle(title string) (film.Film, bool) {
	result := film.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"title": title}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindFilmByID(id domains.ID) (film.Film, bool) {
	result := film.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindPersonByNameAndBirthday(name string, birthday string) (domains.Person, bool) {
	result := domains.Person{}
	err := conn.persons.FindOne(conn.context, bson.M{"name": name, "birthday": birthday}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindPersonByID(id domains.ID) (domains.Person, bool) {
	result := domains.Person{}
	err := conn.persons.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindReviewByID(id domains.ID) (review.Review, bool) {
	result := review.Review{}
	err := conn.reviews.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) GetByQuery(collectionName string, pipeline interface{}) ([]interface{}, *domains.Error) {
	switch collectionName {
	case configs.Default.FilmsCollectionName:
		return AggregateFilms(conn, pipeline)
	case configs.Default.PersonsCollectionName:
		return AggregatePersons(conn, pipeline)
	default:
		return nil, domains.NewError(http.StatusBadRequest, "not supported type")
	}
}

func (conn *MongoConnection) GetFilmsSortedByMark(limit, offset int) ([]film.Film, *domains.Error) {
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

func (conn *MongoConnection) GetFilmsOfGenreSortedByMark(genre domains.Genre, limit, offset int) ([]film.Film, *domains.Error) {
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

func (conn *MongoConnection) GetFilmsOfYearSortedByMark(year, limit, offset int) ([]film.Film, *domains.Error) {
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

func (conn *MongoConnection) GetReviewsSortedByDate(limit, offset int) ([]review.Review, *domains.Error) {
	pipeline := []bson.M{
		{"$sort": bson.M{"date": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	return AggregateReviews(conn, pipeline)
}

func (conn *MongoConnection) GetReviewsOfFilmSortedByDate(filmID domains.ID, limit, offset int) ([]review.Review, *domains.Error) {
	pipeline := []bson.M{
		{"$match": bson.M{"filmid": filmID}},
		{"$sort": bson.M{"date": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	return AggregateReviews(conn, pipeline)
}

func (conn *MongoConnection) GetReviewsOfAuthorSortedByDate(authorID domains.ID, limit, offset int) ([]review.Review, *domains.Error) {
	pipeline := []bson.M{
		{"$match": bson.M{"authorid": authorID}},
		{"$sort": bson.M{"date": -1}},
		{"$limit": limit},
		{"$skip": offset},
	}
	return AggregateReviews(conn, pipeline)
}
