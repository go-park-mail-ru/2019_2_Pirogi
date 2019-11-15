package search

import (
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestStringFromStringArray(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := "[a, b, c]"
	require.Equal(t, expected, stringFromStringArray(input))
}

func TestQuerySearchParams_GetPipelineForMongo(t *testing.T) {
	qp := QuerySearchParams{
		Query:      "Матри",
		Genres:     []string{"драма", "боевик"},
		PersonsIDs: []int{1, 2, 3},
		YearMin:    1950,
		YearMax:    2016,
		Countries:  []string{"США"},
		RatingMin:  0,
		Offset:     5,
		Limit:      10,
		OrderBy:    "",
	}
	expected := []primitive.M{
		{"$limit": 10},
		{"$skip": 5},
		{"$sort": primitive.M{"": -1}},
		{"$match": primitive.M{"year": primitive.M{"$gt": 1950, "$lt": 2016}}},
		{"$match": primitive.M{"genre": primitive.M{"$all": []string{"драма", "боевик"}}}},
		{"$match": primitive.M{"personsid": primitive.M{"$all": []int{1, 2, 3}}}},
		{"$match": primitive.M{"country": primitive.M{"$all": []string{"США"}}}},
		{"$match": primitive.M{"title": primitive.M{"$regex": primitive.Regex{Pattern: ".*Матри.*"}, "$options": "i"}}}}
	require.Equal(t, expected, qp.GetPipelineForMongo("films"))
}