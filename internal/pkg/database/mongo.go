package database

import (
	"context"
	"log"

	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client           *mongo.Client
	users            *mongo.Collection
	usersSize        int
	films            *mongo.Collection
	filmsSize        int
	usersAuthCookies *mongo.Collection
	cookiesSize      int
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

func InitMongo() *MongoConnection {
	client, err := getMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	conn := MongoConnection{
		client:           client,
		users:            client.Database(configs.MongoDbName).Collection(configs.UsersCollectionName),
		usersSize:        0,
		films:            client.Database(configs.MongoDbName).Collection(configs.FilmsCollectionName),
		filmsSize:        0,
		usersAuthCookies: client.Database(configs.MongoDbName).Collection(configs.CoockiesCollectionName),
		cookiesSize:      0,
	}
	return &conn
}

func (conn *MongoConnection) GetID(target string) int {
	switch target {
	case "user":
		return conn.usersSize
	case "film":
		return conn.filmsSize
	case "auth_cookie":
		return conn.cookiesSize
	default:
		return 0
	}
}

func (conn *MongoConnection) Insert(in interface{}) *models.Error {
	switch in := in.(type) {
	case models.NewUser:
		/*_, ok := conn.FindByEmail(in.Email)
		if ok {
			return Error.New(400, "user with the email already exists")
		}*/
		u, e := user.CreateNewUser(conn.GetID("user"), in)
		if e != nil {
			return e
		}
		_, err := conn.users.InsertOne(context.TODO(), u)
		if err != nil {
			return Error.New(500, "cannot insert user in database")
		}
		conn.usersSize++
		return nil
	case models.User:
		filter := bson.D{{"id", in.ID}}
		update := bson.M{
			"$set": bson.M{
				"id":        "ObjectRocket UPDATED!!",
				"fieldbool": false,
			},
		}
		_, err := conn.users.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		return Error.New(404, "user not found")
	/*case models.NewFilm:
		// It is supposed that there cannot be films with the same title
		_, ok := db.FindFilmByTitle(in.Title)
		if ok {
			return Error.New(400, "film with the title already exists")
		}
		f, e := film.CreateNewFilm(db.GetID("film"), &in)
		if e != nil {
			return e
		}
		db.films[db.GetID("film")] = f
		return nil
	case models.Film:
		if _, ok := db.users[in.ID]; ok {
			db.films[in.ID] = in
			return nil
		}
		return Error.New(404, "film not found")*/
	default:
		return Error.New(400, "not supported type")
	}
}

/*func (conn *MongoConnection) FindByEmail(email string) (models.User, bool) {
	for k, u := range db.users {
		if u.Email == email {
			return db.users[k], true
		}
	}
	return models.User{}, false
}*/
