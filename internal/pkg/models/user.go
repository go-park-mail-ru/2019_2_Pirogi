package models

import "net/http"

type Credentials struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"password"`
}

type NewUser struct {
	Credentials `valid:"required"`
	Username    string `json:"username" valid:"stringlength(2|50)"`
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
	Username    string `json:"username,omitempty" valid:"text, stringlength(4|50), optional"`
	Description string `json:"description,omitempty" valid:"description, stringlength(8|50), optional"`
}

type UserCookie struct {
	UserID ID           `json:"user-id" bson:"_id" valid:"numeric"`
	Cookie *http.Cookie `valid:"cookie"`
}
