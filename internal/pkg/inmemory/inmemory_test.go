package inmemory

import (
	"../user"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
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
	u := obj.(user.User)
	require.True(t, reflect.DeepEqual(u, ksyusha))
}
