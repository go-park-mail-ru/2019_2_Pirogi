package user

import (
	"../validators"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type User struct {
	ID         int     `json:"user_id"`
	Email      string  `json:"email"`
	Name       string  `json:"name"`
	Password   string  `json:"password"`
	Rating     float32 `json:"rating"`
	AvatarLink string  `json:"avatar_link"`
}

func CreateUser(id int, email, name, password, avatarLink string, rating float32) (User, error) {
	if !validators.ValidateEmail(email) {
		return User{}, errors.New("email is incorrect")
	}

	// TODO: add validators

	user := User{
		ID:         id,
		Email:      email,
		Name:       name,
		Password:   password,
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
