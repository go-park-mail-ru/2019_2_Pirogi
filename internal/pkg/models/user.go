package models

import "net/http"

type Credentials struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"alphanum, stringlength(8|50)"`
}

type NewUser struct {
	Credentials `valid:"required"`
	Username    string `json:"username" valid:"alphanum, stringlength(4|50)"`
}

type UserTrunc struct {
	ID          ID     `json:"id" valid:"numeric"`
	Username    string `json:"username" valid:"title"`
	Mark        Mark   `json:"mark" valid:"mark, optional"`
	Description string `json:"description" valid:"description"`
	Image       Image  `json:"image" valid:"image, optional"`
}

type User struct {
	Credentials `valid:"required"`
	UserTrunc   `valid:"required"`
}

type UpdateUser struct {
	Username    string `json:"username,omitempty" valid:"username, stringlength(4|50), optional"`
	Description string `json:"description,omitempty" valid:"user_description, stringlength(8|50), optional"`
}

type UserCookie struct {
	UserID ID           `json:"user-id" bson:"_id" valid:"numeric"`
	Cookie *http.Cookie `valid:"cookie"`
}
