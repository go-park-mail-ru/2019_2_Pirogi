package user

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/models"
)

func CreateUser(id domains.ID, newUser *domains.NewUser) (domains.User, *domains.Error) {
	user := domains.User{
		Credentials: models.Credentials{
			Email:    newUser.Email,
			Password: newUser.Password,
		},
		UserTrunc: domains.UserTrunc{
			ID:          id,
			Username:    newUser.Username,
			Mark:        0,
			Description: "",
			Image:       "default.png",
		},
	}
	return user, nil
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
