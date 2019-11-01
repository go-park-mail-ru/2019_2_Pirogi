package models

// struct for film rating
type Rating struct {
	UserID ID   `json:"user_id, omitempty"`
	FilmID ID   `json:"film_id"`
	Mark   Mark `json:"mark"`
}
