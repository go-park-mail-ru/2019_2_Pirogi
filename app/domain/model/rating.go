package model

type RatingNew struct {
	FilmID ID   `json:"film_id" valid:"numeric"`
	Mark   Mark `json:"mark" valid:"mark, optional"`
}

type Rating struct {
	UserID ID   `json:"user_id, omitempty" valid:"numeric"`
	FilmID ID   `json:"film_id" valid:"numeric"`
	Mark   Mark `json:"mark" valid:"mark, optional"`
}

// Чтобы база понимала, обновлять или создавать рейтинг
type RatingUpdate struct {
	UserID ID   `json:"user_id, omitempty" valid:"numeric"`
	FilmID ID   `json:"film_id" valid:"numeric"`
	Mark   Mark `json:"mark" valid:"mark, optional"`
}

func (rn *RatingNew) ToRating(userId ID) Rating {
	return Rating{
		UserID: userId,
		FilmID: rn.FilmID,
		Mark:   rn.Mark,
	}
}

func (s *Rating) SetMark(mark Mark) {
	s.Mark = mark
}