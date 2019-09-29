package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/validators"
)

func CreateUser(id int, email, name, password, avatarLink string, rating float32) (models.User, error) {
	if !validators.ValidateEmail(email) {
		return models.User{}, errors.New("email is incorrect")
	}

	user := models.User{
		ID: id,
		Credentials: models.Credentials{
			Email:    email,
			Password: password,
		},
		Name:       name,
		Rating:     rating,
		AvatarLink: avatarLink,
	}
	return user, nil
}

// temporary function
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
