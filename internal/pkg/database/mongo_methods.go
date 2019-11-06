package database

import (
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNextSequence(layer DatabaseLayer, target string) (models.ID, error) {
	result := struct {
		Seq int `bson:"seq"`
	}{}
	err := layer.FindOneAndUpdateAndDecode(configs.Default.CountersCollectionName, bson.M{"_id": target},
		bson.M{"$inc": bson.M{"seq": 1}}, &result)
	return models.ID(result.Seq), errors.Wrap(err, "get next sequence failed")
}
/*
func InsertUser(conn Database, layer DatabaseLayer, in models.NewUser) *models.Error {
	_, ok := conn.FindUserByEmail(in.Email)
	if ok {
		return Error.New(http.StatusBadRequest, "user with the email already exists")
	}
	id, err := GetNextSequence(layer, configs.Default.UserTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert user in database")
	}
	u, e := user.CreateUser(id, &in)
	if e != nil {
		return e
	}
	err = layer.InsertOne(configs.Default.UsersCollectionName, u)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert user in database")
	}
	return nil
}

func UpdateUser(layer DatabaseLayer, in models.User) *models.Error {
	filter := bson.M{"usertrunc.id": in.ID}
	update := bson.M{"$set": in}
	err := layer.UpdateOne(configs.Default.UsersCollectionName, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "user not found")
	}
	return nil
}

// It is supposed that there cannot be films with the same title
func InsertFilm(conn Database, layer DatabaseLayer, in models.NewFilm) *models.Error {
	_, ok := conn.FindFilmByTitle(in.Title)
	if ok {
		return Error.New(http.StatusBadRequest, "film with the title already exists")
	}
	id, err := GetNextSequence(layer, configs.Default.FilmTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert film in database")
	}
	f := makers.MakeFilm(id, &in)
	err = layer.InsertOne(configs.Default.FilmsCollectionName, f)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert film in database")
	}
	return nil
}

func UpdateFilm(layer DatabaseLayer, in models.Film) *models.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	err := layer.UpdateOne(configs.Default.FilmsCollectionName, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "film not found")
	}
	return nil
}

func UpsertUserCookie(layer DatabaseLayer, in models.UserCookie) *models.Error {
	filter := bson.M{"_id": in.UserID}
	foundCookie := models.UserCookie{}
	err := layer.FindOneAndDecode(configs.Default.CookiesCollectionName, filter, &foundCookie)
	if err != nil {
		err = layer.InsertOne(configs.Default.CookiesCollectionName, in)
	} else {
		update := bson.M{"$set": in}
		err = layer.UpdateOne(configs.Default.CookiesCollectionName, filter, update)
	}
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert cookie in database")
	}
	return nil
}

// It is supposed that there cannot be persons with the same name and birthday
func InsertPerson(conn Database, layer DatabaseLayer, in models.NewPerson) *models.Error {
	_, ok := conn.FindPersonByNameAndBirthday(in.Name, in.Birthday)
	if ok {
		return Error.New(http.StatusBadRequest, "person with this name and birthday already exists")
	}

	id, err := GetNextSequence(layer, configs.Default.PersonTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	newPerson := makers.MakePerson(id, in)
	err = layer.InsertOne(configs.Default.PersonsCollectionName, newPerson)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert person in database: "+err.Error())
	}
	return nil
}

func UpdatePerson(layer DatabaseLayer, in models.Person) *models.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	err := layer.UpdateOne(configs.Default.PersonsCollectionName, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "person not found")
	}
	return nil
}

// TODO: check that user and film of review exist
func InsertReview(layer DatabaseLayer, in models.NewReview) *models.Error {
	id, err := GetNextSequence(layer, configs.Default.ReviewTargetName)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert review in database")
	}
	rev := makers.MakeReview(id, in)
	err = layer.InsertOne(configs.Default.ReviewsCollectionName, rev)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert review in database")
	}
	return nil
}

func UpdateReview(layer DatabaseLayer, in models.Review) *models.Error {
	filter := bson.M{"_id": in.ID}
	update := bson.M{"$set": in}
	err := layer.UpdateOne(configs.Default.ReviewsCollectionName, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "review not found")
	}
	return nil
}

func InsertStars(layer DatabaseLayer, in models.Stars) *models.Error {
	filter := bson.M{"_id": in.FilmID}
	// TODO: рассчитывать newMark
	newMark := in.Mark
	update := bson.M{"$set": bson.M{"mark": newMark}}
	err := layer.UpdateOne(configs.Default.FilmsCollectionName, filter, update)
	if err != nil {
		return Error.New(http.StatusNotFound, "film not found")
	}
	return nil
}

func InsertLike(layer DatabaseLayer, in models.Like) *models.Error {
	err := layer.InsertOne(configs.Default.LikesCollectionName, in)
	if err != nil {
		return Error.New(http.StatusInternalServerError, "cannot insert like in database")
	}
	return nil
}

func AggregateFilms(layer DatabaseLayer, pipeline interface{}) ([]models.Film, *models.Error) {
	result, err := layer.AggregateFilms(configs.Default.FilmsCollectionName, pipeline)
	if err != nil {
		return result, Error.New(http.StatusInternalServerError, err.Error())
	}
	return result, nil
}

func AggregateReviews(layer DatabaseLayer, pipeline interface{}) ([]models.Review, *models.Error) {
	result, err := layer.AggregateReviews(configs.Default.ReviewsCollectionName, pipeline)
	if err != nil {
		return result, Error.New(http.StatusInternalServerError, err.Error())
	}
	return result, nil
}*/
