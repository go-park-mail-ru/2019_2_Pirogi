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
	YearMin    int      `json:"year_min" valid:"year, optional"`
	YearMax    int      `json:"year_max" valid:"year, optional"`
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
		fallthrough
	case qp.OrderBy == "":
		qp.OrderBy = configs.Default.DefaultOrderBy
		fallthrough
	case qp.YearMin == 0:
		qp.YearMin = configs.Default.DefaultYearMin
		fallthrough
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
	switch {
	case qp.YearMin > 0 || qp.YearMax > 0:
		matchBSON = append(matchBSON, bson.M{"year": bson.M{
			"$range": []int{qp.YearMin, qp.YearMax}},
		})
		fallthrough
	case len(qp.Genres) > 0:
		matchBSON = append(matchBSON, bson.M{"genre": stringFromStringArray(qp.Genres)})
		fallthrough
	case len(qp.PersonsIDs) > 0:
		matchBSON = append(matchBSON, bson.M{"personsid": qp.PersonsIDs})
		fallthrough
	case len(qp.Countries) > 0:
		matchBSON = append(matchBSON, bson.M{"country": stringFromStringArray(qp.Countries)})
		fallthrough
	case qp.Query != "":
		switch target {
		case configs.Default.PersonTargetName:
			matchBSON = append(matchBSON, bson.M{"name": queryWrapper(qp.Query)})
		default:
			matchBSON = append(matchBSON, bson.M{"title": queryWrapper(qp.Query)})
		}
	}
	baseBSON = append(baseBSON, matchBSON...)
	return baseBSON
}

func queryWrapper(query string) primitive.M {
	return primitive.M{"$regex": ".*" + query + ".*"}
}

func stringFromStringArray(arr []string) (result string) {
	result = "[" + arr[0]
	for i := 1; i < len(arr); i++ {
		result += ", " + arr[i]
	}
	result += "]"
	return
}
