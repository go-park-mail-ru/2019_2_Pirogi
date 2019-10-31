package database

import (
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	Person "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/person"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
)

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
	filter := bson.M{"_id": in.ID}
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
		return Error.New(http.StatusInternalServerError, "cannot insert user in database")
	}
	f, e := film.CreateFilm(id, &in)
	if e != nil {
		return e
	}
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

func InsertOrUpdateUserCookie(conn *MongoConnection, in models.UserCookie) *models.Error {
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
	return nil
}

func InsertOrUpdatePerson(conn *MongoConnection, in models.Person) *models.Error {
	var (
		person models.Person
		ok     bool
	)
	if in.ID == -1 {
		person, ok = conn.FindPersonByNameAndBirthday(in.Name, in.Birthday)
	} else {
		person, ok = conn.FindPersonByID(in.ID)
	}

	// if we do not have user in database
	if !ok {
		id, err := conn.GetNextSequence(configs.Default.UserTargetName)
		if err != nil {
			return Error.New(http.StatusInternalServerError, "cannot insert user in database")
		}
		newPerson, e := Person.CreatePerson(id, person)
		if e != nil {
			return Error.New(http.StatusInternalServerError, "cannot insert user in database")
		}
		_, err = conn.persons.InsertOne(conn.context, newPerson)
	}
	//TODO: finish the method
	return nil
}

func InsertLike(conn *MongoConnection, in models.Like) *models.Error {
	return nil
}

func InsertReview(conn *MongoConnection, in models.Review) *models.Error {
	return nil
}
