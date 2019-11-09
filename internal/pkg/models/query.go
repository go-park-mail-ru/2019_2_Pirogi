package models

import "go.mongodb.org/mongo-driver/bson"

type QuerySearchParams struct {
	Query     string  `json:"query" valid:"genre"` // TODO: change genre
	Genre     string  `json:"genre" valid:"genre, optional"`
	ActorID   int     `json:"actor_id" valid:"numeric, optional"`
	Year      int     `json:"year" valid:"year, optional"`
	Country   string  `json:"country" valid:"title, optional"` // TODO: change title to country
	RatingMin float32 `json:"rating_min" valid:"mark, optional"`
	Offset    int     `json:"offset" valid:"numeric, optional"`
	Limit     int     `json:"limit" valid:"numeric, optional"`
	SortBy    string  `json:"sort_by" valid:"title, optional"` // TODO: change title
}

func (qp *QuerySearchParams) GetPipelineForMongo() interface{} {
	return []bson.M{
		{"$match": bson.M{
			"genres": qp.Genre,
			"personsid": qp.ActorID,
			"year": qp.Year,
		}},
		{"$sort": bson.M{qp.SortBy: -1}},
		{"$limit": qp.Limit},
		{"$skip": qp.Offset},
	}
}

func (qlp *QueryListParams) getBSON() string {
	return "$limit:" + string(qlp.Limit)
}
