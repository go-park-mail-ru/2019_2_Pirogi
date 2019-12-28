package queryWorker

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/common"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type querySearchParams struct {
	Query      string   `json:"query" valid:"text, optional"`
	Genres     []string `json:"genres" valid:"text, optional"`
	PersonsIds []int    `json:"persons_ids" valid:"numeric, optional"`
	YearMin    int      `json:"year_min" valid:"optional"`
	YearMax    int      `json:"year_max" valid:"optional"`
	Countries  []string `json:"countries" valid:"text, optional"`
	RatingMin  float32  `json:"rating_min" valid:"numeric, optional"`
	Offset     int      `json:"offset" valid:"numeric, optional"`
	Limit      int      `json:"limit" valid:"numeric, optional"`
	OrderBy    string   `json:"order_by" valid:"text, optional"`
}

func NewQueryParams() *querySearchParams {
	return &querySearchParams{}
}

func (qp *querySearchParams) filter() {
	if qp.Limit == 0 {
		qp.Limit = configs.Default.DefaultEntriesLimit
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

func sort(query interface{}) bson.M {
	return bson.M{"$sort": query}
}

func all(query interface{}) bson.M {
	return bson.M{"$all": query}
}

func (qp *querySearchParams) MapQueryParams(ctx echo.Context) {
	qp.Limit = configs.Default.DefaultEntriesLimit // limit must be positive, default value(0) is not suitable
	p := reflect.ValueOf(qp).Elem()
	t := reflect.TypeOf(*qp)
	for i := 0; i < p.NumField(); i++ {
		fieldName := common.ToSnakeCase(t.Field(i).Name)
		switch p.Field(i).Kind() {
		case reflect.Int:
			val, err := strconv.Atoi(ctx.QueryParam(fieldName))
			if err != nil {
				continue
			}
			p.Field(i).SetInt(int64(val))
			continue
		case reflect.Float32:
			val, err := strconv.ParseFloat(ctx.QueryParam(fieldName), 32)
			if err != nil {
				continue
			}
			p.Field(i).SetFloat(val)
			continue
		case reflect.String:
			param := ctx.QueryParam(fieldName)
			if param != "" {
				p.Field(i).SetString(param)
			}
		case reflect.Slice:
			switch t.Field(i).Type.Elem().Kind() {
			case reflect.String:
				querySlice := strings.Split(ctx.QueryParam(fieldName), ",")
				newStringSlice := reflect.MakeSlice(reflect.TypeOf([]string{}), 0, 0)
				for _, item := range querySlice {
					if item != "" {
						newStringSlice = reflect.Append(newStringSlice, reflect.ValueOf(item))
					}
				}
				p.Field(i).Set(newStringSlice)
			case reflect.Int:
				querySlice := strings.Split(ctx.QueryParam(fieldName), ",")
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

func (qp *querySearchParams) GeneratePipeline(target string) []bson.M {
	qp.filter()
	baseBSON := []bson.M{
		{"$limit": qp.Limit},
		{"$skip": qp.Offset},
	}
	var paramsBSON []bson.M
	if qp.YearMin != 0 || qp.YearMax != 0 {
		if qp.YearMin == 0 {
			qp.YearMin = configs.Default.DefaultYearMin
		} else if qp.YearMax == 0 {
			qp.YearMax = configs.Default.DefaultYearMax
		}
		paramsBSON = append(paramsBSON, match(bson.M{"year": bson.M{"$gte": qp.YearMin, "$lte": qp.YearMax}}))
	}
	if len(qp.Genres) > 0 {
		var regexpGenres []primitive.Regex
		for _, genre := range qp.Genres {
			regexpGenres = append(regexpGenres, pattern(genre))
		}
		paramsBSON = append(paramsBSON, match(bson.M{"genres": all(regexpGenres)}))
	}
	if len(qp.PersonsIds) > 0 {
		paramsBSON = append(paramsBSON, match(bson.M{"personsid": all(qp.PersonsIds)}))
	}
	if len(qp.Countries) > 0 {
		var regexp_countries []primitive.Regex
		for _, country := range qp.Countries {
			regexp_countries = append(regexp_countries, pattern(country))
		}
		paramsBSON = append(paramsBSON, match(bson.M{"countries": all(regexp_countries)}))
	}
	if qp.RatingMin != 0 {
		paramsBSON = append(paramsBSON, match(bson.M{"mark": bson.M{"$gte": qp.RatingMin}}))
	}
	if qp.Query != "" {
		switch target {
		case configs.Default.PersonTargetName:
			paramsBSON = append(paramsBSON, match(bson.M{"name": regexp(qp.Query)}))
		default:
			paramsBSON = append(paramsBSON, match(bson.M{"title": regexp(qp.Query)}))
		}
	} else if qp.OrderBy != "" {
		paramsBSON = append(paramsBSON, sort(bson.M{qp.OrderBy: -1}))
	} else if len(paramsBSON) == 0 {
		return nil // В случае запроса без параметров (либо только с limit и orderby)
	}
	paramsBSON = append(paramsBSON, baseBSON...)
	return paramsBSON
}

func GetPipelineForMongoByContext(ctx echo.Context, target string) []bson.M {
	qp := querySearchParams{}
	qp.MapQueryParams(ctx)
	return qp.GeneratePipeline(target)
}

func GetCustomPipelineForMongo(limit, offset int, orderBy, target string) []bson.M {
	qp := querySearchParams{}
	qp.Limit = limit
	qp.Offset = offset
	qp.OrderBy = orderBy
	return qp.GeneratePipeline(target)
}
