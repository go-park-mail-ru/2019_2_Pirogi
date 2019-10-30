package models

import "net/http"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	Username    string  `json:"username"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type User struct {
	ID ID `json:"id" bson:"_id"`
	Credentials
	UserInfo
}

type NewUser struct {
	Credentials
	Username string `json:"username"`
}

type UpdateUser struct {
	Username    string `json:"username,omitempty"`
	Description string `json:"description,omitempty"`
}

type UserCookie struct {
	UserID ID `json:"user-id" bson:"_id"`
	Cookie *http.Cookie
}
