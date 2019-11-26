package model

type StarsRepository interface {
	Insert(stars Stars) (ID, error)
	Update(mark Mark) error
	Delete(id ID) bool
}

type Stars struct {
	UserID ID   `json:"user_id, omitempty" valid:"numeric"`
	FilmID ID   `json:"film_id" valid:"numeric"`
	Mark   Mark `json:"mark" valid:"mark, optional"`
}

func (s *Stars) SetMark(mark Mark) {
	s.Mark = mark
}
