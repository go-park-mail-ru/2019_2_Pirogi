package user

import (
	"../validators"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type User struct {
	Credentials
	ID         int     `json:"user_id"`
	Name       string  `json:"name"`
	Rating     float32 `json:"rating"`
	AvatarLink string  `json:"avatar_link"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(id int, email, name, password, avatarLink string, rating float32) (User, error) {
	if !validators.ValidateEmail(email) {
		return User{}, errors.New("email is incorrect")
	}

	// TODO: add validators

	user := User{
		Credentials: Credentials{
			Email:    email,
			Password: password,
		},
		ID:         id,
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
