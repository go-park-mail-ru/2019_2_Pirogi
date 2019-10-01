package models

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserInfo struct {
	ID          int     `json:"-"`
	Username    string  `json:"username,omitempty"`
	Rating      float32 `json:"rating,omitempty"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"image,omitempty"`
}

type User struct {
	Credentials
	UserInfo
}

type NewUser struct {
	Credentials
	Username string `json:"name"`
}

type ReviewsNum struct {
	Total    int `json:"total"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

type Film struct {
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

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
