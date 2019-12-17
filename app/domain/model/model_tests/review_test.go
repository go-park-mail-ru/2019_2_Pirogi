package model_tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeReview(t *testing.T) {
	expected := testReview
	actual := testReviewNew.ToReview(2)
	expected.Date = actual.Date
	require.Equal(t, expected, actual)
}

