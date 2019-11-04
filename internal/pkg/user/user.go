package user

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/validators"
)

func CreateNewUser(id int, newUser *models.NewUser) (models.User, *models.Error) {
	if !validators.ValidateEmail(newUser.Email) {
		return models.User{}, Error.New(http.StatusBadRequest, "invalid data")
	}
	user := models.User{
		ID: id,
		Credentials: models.Credentials{
			Email:    newUser.Email,
			Password: newUser.Password,
		},
		UserInfo: models.UserInfo{
			Username:    newUser.Username,
			Rating:      0,
			Description: "",
			Image:       "default.jpg",
		},
	}
	return user, nil
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
