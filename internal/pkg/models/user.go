package models

import "net/http"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	Credentials
	Username string `json:"username"`
}

type UserTrunc struct {
	ID          ID      `json:"id"`
	Username    string  `json:"username"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
	Image       Image   `json:"image"`
}

type User struct {
	Credentials
	UserTrunc
}

type UpdateUser struct {
	Username    string `json:"username,omitempty"`
	Description string `json:"description,omitempty"`
}

type UserCookie struct {
	UserID ID `json:"user-id" bson:"_id"`
	Cookie *http.Cookie
}
