package database

import (
	"context"
	"net/http"
	"strconv"

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
	client   *mongo.Client
	context  context.Context
	users    *mongo.Collection
	films    *mongo.Collection
	cookies  *mongo.Collection
	counters *mongo.Collection
}

func getMongoClient() (*mongo.Client, error) {
	credentials := &options.Credential{
		Username: configs.Default.MongoUser,
		Password: configs.Default.MongoPwd,
	}
	clientOpt := &options.ClientOptions{Auth: credentials}
	clientOpt.ApplyURI(configs.Default.MongoHost)
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
		client:   client,
		context:  context.Background(),
		users:    client.Database(configs.Default.MongoDbName).Collection(configs.Default.UsersCollectionName),
		films:    client.Database(configs.Default.MongoDbName).Collection(configs.Default.FilmsCollectionName),
		cookies:  client.Database(configs.Default.MongoDbName).Collection(configs.Default.CookiesCollectionName),
		counters: client.Database(configs.Default.MongoDbName).Collection(configs.Default.CountersCollectionName),
	}

	return &conn, err
}

func (conn *MongoConnection) InitCounters() error {
	_, err := conn.counters.InsertMany(conn.context, []interface{}{
		bson.M{"_id": configs.Default.UserTargetName, "seq": 0},
		bson.M{"_id": configs.Default.FilmTargetName, "seq": 0},
	})
	return err
}

func (conn *MongoConnection) GetNextSequence(target string) (int, error) {
	result := struct {
		Seq int `bson:"seq"`
	}{}
	err := conn.counters.FindOneAndUpdate(conn.context, bson.M{"_id": target},
		bson.M{"$inc": bson.M{"seq": 1}}).Decode(&result)
	return result.Seq, err
}

// затирает старые записи
func (conn *MongoConnection) Insert(in interface{}) *models.Error {
	switch in := in.(type) {
	case models.NewUser:
		_, ok := conn.FindUserByEmail(in.Email)
		if ok {
			return Error.New(http.StatusBadRequest, "user with the email already exists")
		}
		id, err := conn.GetNextSequence(configs.Default.UserTargetName)
		if err != nil {
			return Error.New(http.StatusInternalServerError, "cannot insert user in database")
		}
		u, e := user.CreateNewUser(id, &in)
		if e != nil {
			return e
		}
		_, err = conn.users.InsertOne(conn.context, u)
		if err != nil {
			return Error.New(http.StatusInternalServerError, "cannot insert user in database")
		}
	case models.User:
		filter := bson.M{"_id": in.ID}
		update := bson.M{"$set": in}
		_, err := conn.users.UpdateOne(conn.context, filter, update)
		if err != nil {
			return Error.New(http.StatusNotFound, "user not found")
		}
	case models.NewFilm:
		// It is supposed that there cannot be films with the same title
		_, ok := conn.FindFilmByTitle(in.Title)
		if ok {
			return Error.New(http.StatusBadRequest, "film with the title already exists")
		}
		id, err := conn.GetNextSequence(configs.Default.FilmTargetName)
		if err != nil {
			return Error.New(http.StatusInternalServerError, "cannot insert user in database")
		}
		f, e := film.CreateNewFilm(id, &in)
		if e != nil {
			return e
		}
		_, err = conn.films.InsertOne(conn.context, f)
		if err != nil {
			return Error.New(http.StatusInternalServerError, "cannot insert film in database")
		}
	case models.Film:
		filter := bson.M{"_id": in.ID}
		update := bson.M{"$set": in}
		_, err := conn.films.UpdateOne(conn.context, filter, update)
		if err != nil {
			return Error.New(http.StatusNotFound, "film not found")
		}
	case models.UserCookie:
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
			return Error.New(http.StatusInternalServerError, "cannot insert cookie in database: "+err.Error())
		}
	default:
		return Error.New(http.StatusBadRequest, "not supported type")
	}
	return nil
}

func (conn *MongoConnection) Get(id int, target string) (interface{}, *models.Error) {
	switch target {
	case configs.Default.UserTargetName:
		u, ok := conn.FindUserByID(id)
		if ok {
			return u, nil
		}
		return nil, Error.New(http.StatusNotFound, "no user with id: "+strconv.Itoa(id))
	case configs.Default.FilmTargetName:
		f, ok := conn.FindFilmByID(id)
		if ok {
			return f, nil
		}
		return nil, Error.New(http.StatusNotFound, "no film with the id: "+strconv.Itoa(id))
	}
	return nil, Error.New(http.StatusNotFound, "not supported type: "+target)
}

func (conn *MongoConnection) CheckCookie(cookie *http.Cookie) bool {
	foundCookie := models.UserCookie{}
	err := conn.cookies.FindOne(conn.context, bson.M{"cookie.value": cookie.Value}).Decode(&foundCookie)
	return err == nil
}

func (conn *MongoConnection) DeleteCookie(in interface{}) {
	switch in := in.(type) {
	case http.Cookie:
		u, ok := conn.FindUserByCookie(&in)
		if !ok {
			return
		}
		_, _ = conn.cookies.DeleteOne(conn.context, bson.M{"_id": u.ID})
	}
}

func (conn *MongoConnection) FindUserByEmail(email string) (models.User, bool) {
	result := models.User{}
	err := conn.users.FindOne(conn.context, bson.M{"credentials.email": email}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindUserByID(id int) (models.User, bool) {
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

func (conn *MongoConnection) FindFilmByTitle(title string) (models.Film, bool) {
	result := models.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"filminfo.title": title}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FindFilmByID(id int) (models.Film, bool) {
	result := models.Film{}
	err := conn.films.FindOne(conn.context, bson.M{"_id": id}).Decode(&result)
	return result, err == nil
}

func (conn *MongoConnection) FakeFillDB() {
	cookie := http.Cookie{
		Name:  "cinsear_session",
		Value: "value",
		Path:  "/",
	}

	conn.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "oleg@mail.ru", Password: user.GetMD5Hash("qwerty123")},
		Username:    "Oleg",
	})
	conn.Insert(models.UserCookie{UserID: 0, Cookie: &cookie})

	conn.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "anton@mail.ru", Password: user.GetMD5Hash("qwe523")},
		Username:    "Anton",
	})
	conn.Insert(models.UserCookie{UserID: 1, Cookie: &cookie})

	conn.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "yura@gmail.com", Password: user.GetMD5Hash("12312312")},
		Username:    "Yura",
	})
	conn.Insert(models.UserCookie{UserID: 2, Cookie: &cookie})

	conn.Insert(models.NewFilm{FilmInfo: models.FilmInfo{
		Title: "Бойцовский клуб",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Date:       "1999",
		Actors:     []string{"Брэд Питт", "Эдвард Нортон"},
		Genres:     []string{"Драма", "Боевик"},
		Directors:  []string{"Дэвид Финчер"},
		Rating:     9.1,
		Image:      "club.jpg",
		ReviewsNum: models.ReviewsNum{Total: 100, Positive: 90, Negative: 10},
	}})

	conn.Insert(models.NewFilm{FilmInfo: models.FilmInfo{
		Title: "Матрица",
		Description: "Мир Матрицы — это иллюзия, существующая только в бесконечном сне обреченного человечества. " +
			"Холодный мир будущего, в котором люди — всего лишь батарейки в компьютерных системах.",
		Date:       "1999",
		Actors:     []string{"Киану Ривз", "Кэрри-Энн Мосс"},
		Genres:     []string{"Фэнтези"},
		Directors:  []string{"Лана Вачовски", "Лилли Вачовски"},
		Rating:     8.9,
		Image:      "matrix.jpg",
		ReviewsNum: models.ReviewsNum{Total: 110, Positive: 90, Negative: 20},
	}})
}

func (conn *MongoConnection) ClearDB() {
	_, _ = conn.users.DeleteMany(conn.context, bson.M{})
	_, _ = conn.cookies.DeleteMany(conn.context, bson.M{})
	_, _ = conn.films.DeleteMany(conn.context, bson.M{})
	_, _ = conn.counters.DeleteMany(conn.context, bson.M{})
}
