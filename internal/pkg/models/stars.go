package models

// struct for film rating
type Stars struct {
	UserID ID   `json:"user_id, omitempty" valid:"numeric"`
	FilmID ID   `json:"film_id" valid:"numeric"`
	Mark   Mark `json:"mark" valid:"mark, optional"`
}
