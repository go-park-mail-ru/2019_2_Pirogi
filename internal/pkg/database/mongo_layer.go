package database

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoLayer struct {
	client  *mongo.Client
	context context.Context
	dbName  string
}

func NewMongoLayer(client *mongo.Client, context context.Context, dbName string) *MongoLayer {
	return &MongoLayer{
		client:  client,
		context: context,
		dbName:  dbName,
	}
}

func (m *MongoLayer) collection(collectionName string) *mongo.Collection {
	return m.client.Database(m.dbName).Collection(collectionName)
}

func (m *MongoLayer) FindOneAndUpdateAndDecode(collectionName string, filter interface{},
		update interface{}, v interface{}) error {
	return m.collection(collectionName).FindOneAndUpdate(m.context, filter, update).Decode(v)
}

func (m *MongoLayer) InsertMany(collectionName string, documents []interface{}) error {
	_, err := m.collection(collectionName).InsertMany(m.context, documents)
	return err
}

func (m *MongoLayer) DeleteOne(collectionName string, filter interface{}) error {
	_, err := m.collection(collectionName).DeleteOne(m.context, filter)
	return err
}

func (m *MongoLayer) DeleteMany(collectionName string, filter interface{}) {
	_, _ = m.collection(collectionName).DeleteMany(m.context, filter)
}

func (m *MongoLayer) FindOneAndDecode(collectionName string, filter interface{}, v interface{}) error {
	return m.collection(collectionName).FindOne(m.context, filter).Decode(v)
}

func (m *MongoLayer) InsertOne(collectionName string, document interface{}) error {
	_, err := m.collection(collectionName).InsertOne(m.context, document)
	return err
}

func (m *MongoLayer) UpdateOne(collectionName string, filter interface{}, update interface{}) error {
	_, err := m.collection(collectionName).UpdateOne(m.context, filter, update)
	return err
}

func (m *MongoLayer) AggregateFilms(collectionName string, pipeline interface{}) ([]models.Film, error) {
	curs, err := m.collection(collectionName).Aggregate(m.context, pipeline)
	if err != nil {
		return nil, errors.New("error while aggregating films")
	}
	var result []models.Film
	for curs.Next(m.context) {
		f := models.Film{}
		err = curs.Decode(&f)
		if err != nil {
			return nil, errors.New("error while decoding aggregated result in films")
		}
		result = append(result, f)
	}
	return result, nil
}

func (m *MongoLayer) AggregateReviews(collectionName string, pipeline interface{}) ([]models.Review, error) {
	curs, err := m.collection(collectionName).Aggregate(m.context, pipeline)
	if err != nil {
		return nil, errors.New("error while aggregating reviews")
	}
	var result []models.Review
	for curs.Next(m.context) {
		r := models.Review{}
		err = curs.Decode(&r)
		if err != nil {
			return nil, errors.New("error while decoding aggregated result in reviews")
		}
		result = append(result, r)
	}
	return result, nil
}
