package model

type RatingNew struct {
	FilmID ID    `json:"film_id" valid:"numeric"`
	Stars  Stars `json:"stars" valid:"stars, optional"`
}

type Rating struct {
	UserID ID    `json:"user_id, omitempty" valid:"numeric"`
	FilmID ID    `json:"film_id" valid:"numeric"`
	Stars  Stars `json:"stars" valid:"stars, optional"`
}

// Чтобы база понимала, обновлять или создавать рейтинг
type RatingUpdate struct {
	UserID ID    `json:"user_id, omitempty" valid:"numeric"`
	FilmID ID    `json:"film_id" valid:"numeric"`
	Stars  Stars `json:"stars" valid:"stars, optional"`
}

func (rn *RatingNew) ToRating(userId ID) Rating {
	return Rating{
		UserID: userId,
		FilmID: rn.FilmID,
		Stars:  rn.Stars,
	}
}

func (rn *Rating) ToRatingUpdate() RatingUpdate {
	return RatingUpdate{
		UserID: rn.UserID,
		FilmID: rn.FilmID,
		Stars:  rn.Stars,
	}
}
