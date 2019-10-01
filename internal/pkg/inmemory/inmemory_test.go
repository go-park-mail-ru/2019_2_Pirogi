package inmemory

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFakeFillDB(t *testing.T) {
	db := Init()
	db.FakeFillDB()
	require.Equal(t, 3, len(db.users))
}

func TestDB_Insert(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	err := db.Insert(ksyusha)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(db.users[0].Credentials, ksyusha.Credentials))
}

func TestDB_Get(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	err := db.Insert(ksyusha)
	require.Nil(t, err)
	obj, err := db.Get(0, "user")
	require.Nil(t, err)
	u := obj.(models.User)
	require.True(t, reflect.DeepEqual(u.Credentials, ksyusha.Credentials))
}
