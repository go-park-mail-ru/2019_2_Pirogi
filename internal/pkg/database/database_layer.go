package database

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

type DatabaseLayer interface {
	FindOneAndUpdateAndDecode(collectionName string, filter interface{}, update interface{}, v interface{}) error
	InsertMany(collectionName string, documents []interface{}) error
	DeleteOne(collectionName string, filter interface{}) error
	DeleteMany(collectionName string, filter interface{})
	FindOneAndDecode(collectionName string, filter interface{}, v interface{}) error
	InsertOne(collectionName string, document interface{}) error
	UpdateOne(collectionName string, filter interface{}, update interface{}) error
	AggregateFilms(collectionName string, pipeline interface{}) ([]models.Film, error)
	AggregateReviews(collectionName string, pipeline interface{}) ([]models.Review, error)
}
