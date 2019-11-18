package domains

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
	"testing"
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
	model := ReviewNew{}
	err = model.Make(body)
	require.NoError(t, err)
	require.Equal(t, testReviewNew, model)
}
