package model_tests

import (
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestMakeReview(t *testing.T) {
	expected := testReview
	actual := testReviewNew.ToReview(2)
	expected.Date = actual.Date
	require.Equal(t, expected, actual)
}

func TestPrepareModelNewReview(t *testing.T) {
	validation.InitValidator()
	body, err := testReviewNew.MarshalJSON()
	require.NoError(t, err)
	reviewNew := model.ReviewNew{}
	err = reviewNew.Make(body)
	require.NoError(t, err)
	require.Equal(t, testReviewNew, reviewNew)
}
