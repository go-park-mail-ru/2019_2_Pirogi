package user

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateUser(id models.ID, newUser *models.NewUser) (models.User, *models.Error) {
	user := models.User{
		Credentials: models.Credentials{
			Email:    newUser.Email,
			Password: newUser.Password,
		},
		UserTrunc: models.UserTrunc{
			ID:          id,
			Username:    newUser.Username,
			Mark:        0,
			Description: "",
			Image:       "default.png",
			Reviews:     0,
		},
	}
	return user, nil
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
