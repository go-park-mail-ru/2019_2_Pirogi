package models

import "net/http"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserInfo struct {
	Username    string  `json:"username"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type User struct {
	ID int `json:"id" bson:"_id"`
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

type ReviewsNum struct {
	Total    int `json:"total"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

type FilmInfo struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Actors      []string `json:"actors"`
	Genres      []string `json:"genres"`
	Directors   []string `json:"directors"`
	Rating      float32  `json:"rating"`
	Image       string   `json:"image"`
	ReviewsNum
}

type NewFilm struct {
	FilmInfo
}

type Film struct {
	ID int `json:"id" bson:"_id"`
	FilmInfo
}

type UserCookie struct {
	UserID int `json:"userid" bson:"_id"`
	Cookie *http.Cookie
}

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
