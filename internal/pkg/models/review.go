package models

type ReviewsNum int

type NewReview struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	FilmID   ID       `json:"film_id"`
	AuthorID ID       `json:"author_id"`
	FilmMark FilmMark `json:"film_mark"`
}

// TODO: remove binary choice of film's like/dislike
type Review struct {
	ID    ID     `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	//TODO: заменить на гошный формат
	Date     string   `json:"date"`
	FilmID   ID       `json:"film_id"`
	AuthorID ID       `json:"author_id"`
	FilmMark FilmMark `json:"film_mark"`
	Likes    int      `json:"likes"`
}
