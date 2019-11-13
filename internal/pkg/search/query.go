package search

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuerySearchParams struct {
	Query      string   `json:"query" valid:"text"`
	Genres     []string `json:"genres" valid:"text, optional"`
	PersonsIDs []int    `json:"persons_id" valid:"numeric, optional"`
	YearMin    int      `json:"year_min" valid:"optional"`
	YearMax    int      `json:"year_max" valid:"optional"`
	Countries  []string `json:"countries" valid:"text, optional"`
	RatingMin  float32  `json:"rating_min" valid:"numeric, optional"`
	Offset     int      `json:"offset" valid:"numeric, optional"`
	Limit      int      `json:"limit" valid:"numeric, optional"`
	OrderBy    string   `json:"order_by" valid:"text, optional"`
}

func (qp *QuerySearchParams) filter() {
	switch {
	case qp.Limit == 0:
		qp.Limit = configs.Default.DefaultEntriesLimit
	case qp.OrderBy == "":
		qp.OrderBy = configs.Default.DefaultOrderBy
	case qp.YearMin == 0:
		print(qp.YearMin)
		qp.YearMin = configs.Default.DefaultYearMin
	case qp.YearMax == 0:
		qp.YearMax = configs.Default.DefaultYearMax
	}
}

func (qp *QuerySearchParams) GetPipelineForMongo(target string) interface{} {
	qp.filter()
	baseBSON := []bson.M{
		{"$limit": qp.Limit},
		{"$skip": qp.Offset},
		{"$sort": bson.M{qp.OrderBy: -1}},
	}
	var matchBSON []bson.M
	matchBSON = append(matchBSON, match(primitive.M{"year": primitive.M{"$gt": qp.YearMin,
		"$lt": qp.YearMax},
	}))
	switch {
	case len(qp.Genres) > 0:
		matchBSON = append(matchBSON, match(bson.M{"genre": all(qp.Genres)}))
		fallthrough
	case len(qp.PersonsIDs) > 0:
		matchBSON = append(matchBSON, match(bson.M{"personsid": all(qp.PersonsIDs)}))
		fallthrough
	case len(qp.Countries) > 0:
		matchBSON = append(matchBSON, match(bson.M{"country": all(qp.Countries)}))
		fallthrough
	case qp.Query != "":
		switch target {
		case configs.Default.PersonTargetName:
			matchBSON = append(matchBSON, match(bson.M{"name": regexp(qp.Query)}))
		default:
			matchBSON = append(matchBSON, match(bson.M{"title": regexp(qp.Query)}))
		}
	}
	baseBSON = append(baseBSON, matchBSON...)
	return baseBSON
}

func regexp(query string) primitive.M {
	return bson.M{"$regex": pattern(query), "$options": 'i'}
}

func pattern(query string) primitive.Regex {
	return primitive.Regex{Pattern: ".*" + query + ".*"}
}

func match(query interface{}) primitive.M {
	return primitive.M{"$match": query}
}

func all(query interface{}) primitive.M {
	return primitive.M{"$all": query}
}

func stringFromStringArray(arr []string) (result string) {
	result = "[" + arr[0]
	for i := 1; i < len(arr); i++ {
		result += ", " + arr[i]
	}
	result += "]"
	return
}
