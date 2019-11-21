package queryWorker

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strconv"
	"strings"
)

type querySearchParams struct {
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

func (qp *querySearchParams) filter() {
	if qp.Limit == 0 {
		qp.Limit = configs.Default.DefaultEntriesLimit
	}
	if qp.OrderBy == "" {
		qp.OrderBy = configs.Default.DefaultOrderBy
	}
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

func (qp *querySearchParams) mapQueryParams(ctx echo.Context) {
	qp.Limit = configs.Default.DefaultEntriesLimit // limit must be positive, default value(0) is not suitable
	p := reflect.ValueOf(qp).Elem()
	t := reflect.TypeOf(*qp)
	for i := 0; i < p.NumField(); i++ {
		switch p.Field(i).Kind() {
		case reflect.Int:
			val, err := strconv.Atoi(ctx.QueryParam(strings.ToLower(t.Field(i).Name)))
			if err != nil {
				continue
			}
			p.Field(i).SetInt(int64(val))
			continue
		case reflect.String:
			param := ctx.QueryParam(strings.ToLower(t.Field(i).Name))
			if param != "" {
				p.Field(i).SetString(param)
			}
		case reflect.Slice:
			switch t.Field(i).Type.Elem().Kind() {
			case reflect.String:
				querySlice := strings.Split(ctx.QueryParam(strings.ToLower(t.Field(i).Name)), ",")
				newStringSlice := reflect.MakeSlice(reflect.TypeOf([]string{}), 0, 0)
				for _, item := range querySlice {
					if item != "" {
						newStringSlice = reflect.Append(newStringSlice, reflect.ValueOf(item))
					}
				}
				p.Field(i).Set(newStringSlice)
			case reflect.Int:
				querySlice := strings.Split(ctx.QueryParam(strings.ToLower(t.Field(i).Name)), ",")
				println(strings.ToLower(t.Field(i).Name))
				var newIntValues []int
				for _, item := range querySlice {
					value, err := strconv.Atoi(item)
					if err != nil {
						continue
					}
					newIntValues = append(newIntValues, value)
				}
				newIntSlice := reflect.MakeSlice(reflect.TypeOf([]int{}), 0, 0)
				for _, item := range newIntValues {
					newIntSlice = reflect.Append(newIntSlice, reflect.ValueOf(item))
				}
				p.Field(i).Set(newIntSlice)
			}
		}
	}
}

func (qp *querySearchParams) generatePipeline(target string) interface{} {
	qp.filter()
	baseBSON := []bson.M{
		{"$limit": qp.Limit},
		{"$skip": qp.Offset},
		{"$sort": bson.M{qp.OrderBy: -1}},
	}
	var matchBSON []bson.M
	if qp.YearMin != 0 || qp.YearMax != 0 {
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

func GetPipelineForMongo(ctx echo.Context, target string) interface{} {
	qp := querySearchParams{}
	qp.mapQueryParams(ctx)
	return qp.generatePipeline(target)
}
