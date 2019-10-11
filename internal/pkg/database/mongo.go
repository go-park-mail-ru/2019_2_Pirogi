package database

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client           *mongo.Client
	context          context.Context
	users            *mongo.Collection
	films            *mongo.Collection
	usersAuthCookies *mongo.Collection
	counters         *mongo.Collection
}

func getMongoClient() (*mongo.Client, error) {
	credentials := &options.Credential{
		Username: configs.MongoUser,
		Password: configs.MongoPwd,
	}
	clientOpt := &options.ClientOptions{Auth: credentials}
	clientOpt.ApplyURI(configs.MongoHost)
	client, err := mongo.NewClient(clientOpt)
	return client, err
}

func InitMongo() (*MongoConnection, error) {
	client, err := getMongoClient()
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
		client:           client,
		context:          context.Background(),
		users:            client.Database(configs.MongoDbName).Collection(configs.UsersCollectionName),
		films:            client.Database(configs.MongoDbName).Collection(configs.FilmsCollectionName),
		usersAuthCookies: client.Database(configs.MongoDbName).Collection(configs.CoockiesCollectionName),
		counters:         client.Database(configs.MongoDbName).Collection(configs.CountersCollectionName),
	}

	// Do it one time
	_, _ = conn.counters.InsertMany(conn.context, []interface{}{
		bson.M{"_id": configs.UserTargetName, "seq": 0},
		bson.M{"_id": configs.FilmTargetName, "seq": 0},
		bson.M{"_id": configs.CookieTargetName, "seq": 0},
	})

	return &conn, err
}

func (conn *MongoConnection) GetNextSequence(target string) (int, error) {
	result := struct {
		Seq    int    `bson:"seq"`
	}{}
	err := conn.counters.FindOneAndUpdate(conn.context, bson.M{"_id": target},
		bson.M{"$inc": bson.M{"seq": 1}}).Decode(&result)
	return result.Seq, err
}

func (conn *MongoConnection) Insert(in interface{}) *models.Error {
	switch in := in.(type) {
	case models.NewUser:
		_, ok := conn.FindByEmail(in.Email)
		if ok {
			return Error.New(400, "user with the email already exists")
		}
		id, err := conn.GetNextSequence(configs.UserTargetName)
		if err != nil {
			return Error.New(500, "cannot insert user in database")
		}
		u, e := user.CreateNewUser(id, in)
		if e != nil {
			return e
		}
		_, err = conn.users.InsertOne(conn.context, u)
		if err != nil {
			return Error.New(500, "cannot insert user in database")
		}
		return nil
	case models.User:
		filter := bson.M{"_id": in.ID}
		update := bson.M{"$set": in}
		_, err := conn.users.UpdateOne(conn.context, filter, update)
		if err != nil {
			return Error.New(404, "user not found")
		}
		return nil
	case models.NewFilm:
		// It is supposed that there cannot be films with the same title
		_, ok := conn.FindFilmByTitle(in.Title)
		if ok {
			return Error.New(400, "film with the title already exists")
		}
		id, err := conn.GetNextSequence(configs.FilmTargetName)
		if err != nil {
			return Error.New(500, "cannot insert user in database")
		}
		f, e := film.CreateNewFilm(id, &in)
		if e != nil {
			return e
		}
		_, err = conn.films.InsertOne(conn.context, f)
		if err != nil {
			return Error.New(500, "cannot insert film in database")
		}
		return nil
	case models.Film:
		filter := bson.M{"_id": in.ID}
		update := bson.M{"$set": in}
		_, err := conn.films.UpdateOne(conn.context, filter, update)
		if err != nil {
			return Error.New(404, "film not found")
		}
		return nil
	default:
		return Error.New(400, "not supported type")
	}
}

func (conn *MongoConnection) FindByEmail(email string) (models.User, bool) {
	result := models.User{}
	err := conn.users.FindOne(conn.context, bson.M{"credentials.email": email}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindFilmByTitle(title string) (models.Film, bool) {
	result := models.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"filminfo.title": title}).Decode(&result)
	return result, err == nil
}
