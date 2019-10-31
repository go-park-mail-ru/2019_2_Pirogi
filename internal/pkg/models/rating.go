package models

// struct for film rating
type Rating struct {
	ID     ID   `json:"id, omitempty"`
	UserID ID   `json:"user_id, omitempty"`
	FilmID ID   `json:"film_id"`
	Mark   Mark `json:"mark"`
}
