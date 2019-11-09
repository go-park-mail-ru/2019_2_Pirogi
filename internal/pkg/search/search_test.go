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
		Query:      "матри",
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
		{"year": primitive.M{"$range": []int{1950, 2016}}},
		{"genre": "[драма, боевик]"},
		{"personsid": []int{1, 2, 3}},
		{"country": "[США]"},
		{"title": primitive.M{"$regex": ".*матри.*"}}}
	require.Equal(t, expected, qp.GetPipelineForMongo("films"))
}
