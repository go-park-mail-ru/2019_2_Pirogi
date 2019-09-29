package models

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Credentials
	ID          int     `json:"user_id"`
	Name        string  `json:"name"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
	AvatarLink  string  `json:"avatar_link"`
}

type NewUser struct {
	Credentials
	Name string `json:"name"`
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
	ReviewsNum
	ImagePath string `json:"image_path"`
}

type UpdateUser struct {
	Credentials
	Name        string `json:"name"`
	Description string `json:"description"`
	ActualEmail string `json:"actual_email"`
}

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
