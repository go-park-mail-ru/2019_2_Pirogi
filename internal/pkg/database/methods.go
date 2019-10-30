package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
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
	u, e := user.CreateNewUser(id, &in)
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
func InsertFilm(conn *MongoConnection, in models.NewFilm) *models.Error {
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