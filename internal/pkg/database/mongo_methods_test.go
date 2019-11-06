package database

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	mockdatabase "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database/mock"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type TestCaseGetNextSequence struct {
	Target string
	Result struct {
		Seq int `bson:"seq"`
	}
	GivenError    error
	ExpectedID    models.ID
}

func TestGetNextSequence(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockDbLayer := mockdatabase.NewMockDatabaseLayer(ctrl)
	testCases := []TestCaseGetNextSequence{
		{
			Target:     configs.Default.UserTargetName,
			Result:     struct{ Seq int `bson:"seq"` }{1},
			ExpectedID: 1,
		},
		{
			Target:        "invalidTarget",
			GivenError:    errors.New("some error"),
		},
	}

	for _, testCase := range testCases {
		result := struct {
			Seq int `bson:"seq"`
		}{}
		mockDbLayer.EXPECT().FindOneAndUpdateAndDecode(configs.Default.CountersCollectionName,
			bson.M{"_id": testCase.Target}, bson.M{"$inc": bson.M{"seq": 1}}, &result).SetArg(3, testCase.Result).
			Return(testCase.GivenError)

		id, err := GetNextSequence(mockDbLayer, testCase.Target)
		if testCase.GivenError != nil {
			assert.Error(t, err)
			assert.Equal(t, "get next sequence failed: "+testCase.GivenError.Error(), err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectedID, id)
		}
	}
}

func TestInitCounters(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockDb := mockdatabase.NewMockDatabase(ctrl)
	mockDbLayer := mockdatabase.NewMockDatabaseLayer(ctrl)
	result := models.User{}
	mockDbLayer.EXPECT().FindOneAndDecode(configs.Default.UsersCollectionName, bson.M{"usertrunc.id": 1}, &result).
		Return(nil)
	u, ok := mockDb.FindUserByID(1)
	assert.Equal(t, 1, u.ID)
	assert.Equal(t, true, ok)
}
