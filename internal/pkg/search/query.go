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
	Year       int      `json:"year" valid:"numeric, optional"`
}

func (qp *QuerySearchParams) filter() {
	if qp.Limit == 0 {
		qp.Limit = configs.Default.DefaultEntriesLimit
	}
	if qp.OrderBy == "" {
		qp.OrderBy = configs.Default.DefaultOrderBy
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
	if qp.YearMin != 0 || qp.YearMax != 0  {
		if qp.YearMin == 0 {
			qp.YearMin = configs.Default.DefaultYearMin
		} else if qp.YearMax == 0 {
			qp.YearMax = configs.Default.DefaultYearMax
		}
		matchBSON = append(matchBSON, match(bson.M{"year": bson.M{"$gte": qp.YearMin, "$lte": qp.YearMax}}))
	}
	if len(qp.Genres) > 0 {
		var regexp_genres []primitive.Regex
		for _, genre := range qp.Genres {
			regexp_genres = append(regexp_genres, pattern(genre))
		}
		matchBSON = append(matchBSON, match(bson.M{"genres": all(regexp_genres)}))
	}
	if len(qp.PersonsIDs) > 0 {
		matchBSON = append(matchBSON, match(bson.M{"personsid": all(qp.PersonsIDs)}))
	}
	if len(qp.Countries) > 0 {
		var regexp_countries []primitive.Regex
		for _, country := range qp.Countries {
			regexp_countries = append(regexp_countries, pattern(country))
		}
		matchBSON = append(matchBSON, match(bson.M{"countries": all(regexp_countries)}))
	}
	if qp.Year != 0 {
		matchBSON = append(matchBSON, match(bson.M{"year": qp.Year}))
	}
	if qp.Query != "" {
		switch target {
		case configs.Default.PersonTargetName:
			matchBSON = append(matchBSON, match(bson.M{"name": regexp(qp.Query)}))
		default:
			matchBSON = append(matchBSON, match(bson.M{"title": regexp(qp.Query)}))
		}
	}
	matchBSON = append(matchBSON, baseBSON...)
	return matchBSON
}

func regexp(query string) bson.M {
	return bson.M{"$regex": pattern(query)}
}

func pattern(query string) primitive.Regex {
	return primitive.Regex{Pattern: ".*" + query + ".*", Options: "i"}
}

func match(query interface{}) bson.M {
	return bson.M{"$match": query}
}

func all(query interface{}) bson.M {
	return bson.M{"$all": query}
}

func stringFromStringArray(arr []string) (result string) {
	result = "[" + arr[0]
	for i := 1; i < len(arr); i++ {
		result += ", " + arr[i]
	}
	result += "]"
	return
}
