package inmemory

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/stretchr/testify/require"
)

func TestFakeFillDB(t *testing.T) {
	db := Init()
	db.FakeFillDB()
	require.Equal(t, 3, len(db.users))
}

func TestDB_Insert(t *testing.T) {
	db := Init()
	ksyusha, _ := user.CreateUser(5, "ksyushag@mail.ru", "Ksyusha", user.GetMD5Hash("qwerty123"), "ksyuha.jpg", 8.3)
	err := db.Insert(ksyusha)
	require.NoError(t, err)
	require.True(t, reflect.DeepEqual(db.users[5], ksyusha))
}

func TestDB_Get(t *testing.T) {
	db := Init()
	ksyusha, _ := user.CreateUser(5, "ksyushag@mail.ru", "Ksyusha", user.GetMD5Hash("qwerty123"), "ksyuha.jpg", 8.3)
	err := db.Insert(ksyusha)
	require.NoError(t, err)
	obj, err := db.Get(5, "user")
	require.NoError(t, err)
	u := obj.(models.User)
	require.True(t, reflect.DeepEqual(u, ksyusha))
}
