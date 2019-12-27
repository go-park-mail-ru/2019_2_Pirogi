package model

import (
	"github.com/asaskevich/govalidator"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

type UserNew struct {
	Username string `json:"username" valid:"title"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"password"`
}

type UserCredentials struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"password"`
}

func (un *UserNew) Make(body []byte) error {
	err := un.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(un)
	return err
}

type UserTrunc struct {
	ID          ID     `json:"id" valid:"numeric"`
	Username    string `json:"username" valid:"title"`
	Description string `json:"description" valid:"description"`
	Image       Image  `json:"image" valid:"image, optional"`
}

type User struct {
	ID          ID     `json:"id" valid:"numeric,optional"`
	Email       string `json:"email" valid:"email,optional"`
	Password    string `json:"password,omitempty" valid:"password,optional"`
	Username    string `json:"username" valid:"title,optional"`
	Description string `json:"description" valid:"description,optional"`
	Image       Image  `json:"image" valid:"image,optional"`
}

func (u *User) IsPasswordCorrect(password string) bool {
	return u.Password == password
}

func (un *UserNew) ToUser(id ID) User {
	return User{
		ID:          id,
		Email:       un.Email,
		Password:    un.Password,
		Description: "",
		Image:       Image(configs.Default.DefaultImageName),
	}
}

func (u *User) Trunc() UserTrunc {
	return UserTrunc{
		ID:          u.ID,
		Username:    u.Username,
		Description: u.Description,
		Image:       u.Image,
	}
}

func (u *User) FromProtobuf(user v1.User) {
	u.ID = ID(user.ID)
	u.Email = user.Email
	u.Username = user.Username
	u.Password = user.Password
	u.Description = user.Description
	u.Image = Image(user.Image)
}

func (u *User) ToProtobuf() (user v1.User) {
	return v1.User{
		ID:          int64(u.ID),
		Email:       u.Email,
		Password:    u.Password,
		Username:    u.Username,
		Description: u.Description,
		Image:       string(u.Image),
	}
}

func (u *UserNew) ToProtobuf() (userNew v1.UserNew) {
	return v1.UserNew{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *UserNew) FromProtobuf(userNew v1.UserNew) {
	u.Email = userNew.Email
	u.Password = userNew.Password
}
