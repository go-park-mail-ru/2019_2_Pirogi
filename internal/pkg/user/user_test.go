package user

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestCreateNewUserOK(t *testing.T) {
//	const id = 0
//	newCredentials := models.Credentials{
//		Email:    "aasdasd@aqweqwe.aas",
//		Password: "qwerty123",
//	}
//	newUser := models.NewUser{
//		Credentials: newCredentials,
//		Username:    "qwerty",
//	}
//	newUserInfo := models.UserInfo{
//		Username:    "qwerty",
//		Rating:      0,
//		Description: "",
//		Image:       "default.jpg",
//	}
//	expectedUser := models.User{
//		ID:          0,
//		Credentials: newCredentials,
//		UserInfo:    newUserInfo,
//	}
//	actualUser, e := CreateUser(id, &newUser)
//	require.Nil(t, e)
//	require.Equal(t, expectedUser, actualUser)
//}

func TestCreateNewUserFail(t *testing.T) {
	const id = 0
	newCredentials := models.Credentials{
		Email:    "123",
		Password: "qwerty123",
	}
	newUser := models.NewUser{
		Credentials: newCredentials,
		Username:    "qwerty",
	}
	_, e := CreateUser(id, &newUser)
	require.NotNil(t, e)
}

func TestGetMD5Hash(t *testing.T) {
	const text = "123"
	hasher := md5.New()
	hasher.Write([]byte(text))
	actual := GetMD5Hash(text)
	expected := hex.EncodeToString(hasher.Sum(nil))
	require.Equal(t, actual, expected)
}
