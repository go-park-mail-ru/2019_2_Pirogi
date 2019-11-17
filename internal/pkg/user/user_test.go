package user

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/models"
	"github.com/stretchr/testify/require"
)

func TestCreateNewUser(t *testing.T) {
	const id = 0
	newCredentials := models.Credentials{
		Email:    "123",
		Password: "qwerty123",
	}
	newUser := domains.NewUser{
		Credentials: newCredentials,
		Username:    "qwerty",
	}
	_, e := CreateUser(id, &newUser)
	require.Nil(t, e)
}

func TestGetMD5Hash(t *testing.T) {
	const text = "123"
	hasher := md5.New()
	hasher.Write([]byte(text))
	actual := GetMD5Hash(text)
	expected := hex.EncodeToString(hasher.Sum(nil))
	require.Equal(t, actual, expected)
}
