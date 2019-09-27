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
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
}

type Image struct {
	ID     int    `json:"id"`
	Target string `json:"target"`
	Path   string `json:"path"`
}

type UpdateUser struct {
	Credentials
	Name        string `json:"name"`
	Description string `json:"description"`
	ActualEmail string `json:"actual_email"`
}
