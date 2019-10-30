package models

type ReviewsNum struct {
	Total    int `json:"total"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

func NewReviewsNum() ReviewsNum {
	return ReviewsNum{
		Total:    0,
		Positive: 0,
		Negative: 0,
	}
}

type NewReview struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	FilmID   ID       `json:"film_id"`
	AuthorID ID       `json:"author_id"`
	FilmMark FilmMark `json:"film_mark"`
}

// TODO: remove binary choice of film's like/dislike
type Review struct {
	ID       ID        `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Date     string    `json:"date"`
	FilmID   ID        `json:"film_id"`
	Author   UserTrunc `json:"author"`
	FilmMark FilmMark  `json:"film_mark"`
	Likes    int       `json:"likes"`
}
