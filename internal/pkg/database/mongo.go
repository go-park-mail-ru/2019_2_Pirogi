package database

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	})
	return errors.Wrap(err, "init counters collection failed")
}

func (conn *MongoConnection) InsertOrUpdate(in interface{}) *models.Error {
	var e *models.Error
	switch in := in.(type) {
	case models.NewUser:
		e = InsertUser(conn, in)
	case models.User:
		e = UpdateUser(conn, in)
	case models.NewFilm:
		e = InsertFilm(conn, in)
	case models.Film:
		e = UpdateFilm(conn, in)
	case models.UserCookie:
		e = InsertOrUpdateUserCookie(conn, in)
	case models.Person:
		e = InsertOrUpdatePerson(conn, in)
	case models.Review:
		e = InsertReview(conn, in)
	case models.Like:
		e = InsertLike(conn, in)
	default:
		e = Error.New(http.StatusBadRequest, "not supported type")
	}
	return e
}

func (conn *MongoConnection) Get(id models.ID, target string) (interface{}, *models.Error) {
	switch target {
	case configs.Default.UserTargetName:
		u, ok := conn.FindUserByID(id)
		if ok {
			return u, nil
		}
		return nil, Error.New(http.StatusNotFound, "no user with id: "+id.String())
	case configs.Default.FilmTargetName:
		f, ok := conn.FindFilmByID(id)
		if ok {
			return f, nil
		}
		return nil, Error.New(http.StatusNotFound, "no film with the id: "+id.String())
	}
	return nil, Error.New(http.StatusNotFound, "not supported type: "+target)
}

func (conn *MongoConnection) Delete(in interface{}) *models.Error {
	switch in := in.(type) {
	case http.Cookie:
		u, ok := conn.FindUserByCookie(&in)
		if !ok {
			return nil
		}
		_, err := conn.cookies.DeleteOne(conn.context, bson.M{"_id": u.ID})
		if err != nil {
			return Error.New(http.StatusInternalServerError, "cannot delete cookie from database: "+err.Error())
		}
		return nil
	}
	return Error.New(http.StatusBadRequest, "not supported type")
}

func (conn *MongoConnection) ClearDB() {
	_, _ = conn.users.DeleteMany(conn.context, bson.M{})
	_, _ = conn.cookies.DeleteMany(conn.context, bson.M{})
	_, _ = conn.films.DeleteMany(conn.context, bson.M{})
	_, _ = conn.counters.DeleteMany(conn.context, bson.M{})
}

func (conn *MongoConnection) CheckCookie(cookie *http.Cookie) bool {
	foundCookie := models.UserCookie{}
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	return err == nil
}

func (conn *MongoConnection) FindUserByEmail(email string) (models.User, bool) {
	result := models.User{}
	err := conn.users.FindOne(conn.context, bson.M{"credentials.email": email}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindUserByID(id models.ID) (models.User, bool) {
	result := models.User{}
	err := conn.users.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindUserByCookie(cookie *http.Cookie) (models.User, bool) {
	foundCookie := models.UserCookie{}
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	if err != nil {
		return models.User{}, false
	}
	return conn.FindUserByID(foundCookie.UserID)
}

// TODO
func (conn *MongoConnection) FindUsersByIDs(ids []models.ID) ([]models.User, bool) {
	return nil, false
}

func (conn *MongoConnection) FindFilmByTitle(title string) (models.Film, bool) {
	result := models.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"filminfo.title": title}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindFilmByID(id models.ID) (models.Film, bool) {
	result := models.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"filmtrunc.id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindPersonByNameAndBirthday(name string, birthday string) (models.Person, bool) {
	result := models.Person{}
	err := conn.persons.FindOne(conn.context, bson.M{"person.name": name, "person.birthday": birthday}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindPersonByID(id models.ID) (models.Person, bool) {
	result := models.Person{}
	err := conn.persons.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

// TODO
func (conn *MongoConnection) GetFilmsSortedByRating(limit int, offset int) ([]models.Film, *models.Error) {
	return nil, nil
}

func (conn *MongoConnection) GetFilmsOfGenreSortedByRating(genre models.Genre, limit int, offset int) ([]models.Film, *models.Error) {
	return nil, nil
}

func (conn *MongoConnection) GetFilmsOfYearSortedByRating(year int, limit int, offset int) ([]models.Film, *models.Error) {
	return nil, nil
}


func (conn *MongoConnection) GetReviewsSortedByDate(limit int, offset int) ([]models.Review, *models.Error) {
	return nil, nil
}

func (conn *MongoConnection) GetReviewsOfFilmSortedByDate(filmTitle string, limit int, offset int) ([]models.Review, *models.Error) {
	return nil, nil
}
