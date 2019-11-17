package domains

import "time"

type ReviewRepository interface {
	Insert(newReview NewReview) (ID, error)
	Update(id ID, review Review) error
	Delete(id ID) bool
	Get(id ID) Review
	GetMany(id []ID) []Review
}

type NewReview struct {
	Title  string `json:"title" valid:"title, stringlength(2|50)"`
	Body   string `json:"body" valid:"description, stringlength(8|50)"`
	FilmID ID     `json:"film_id" valid:"numeric"`
	//TODO: убрать отсюда автор ID
	AuthorID ID `json:"author_id, omitempty" valid:"numeric, optional"`
}

type Review struct {
	ID       ID        `json:"id, omitempty" bson:"_id" valid:"required"`
	Title    string    `json:"title" valid:"title, stringlength(2|50)"`
	Body     string    `json:"body" valid:"description"`
	FilmID   ID        `json:"film_id" valid:"numeric"`
	AuthorID ID        `json:"author_id, omitempty" valid:"numeric"`
	Date     time.Time `json:"date" valid:"time"`
	Likes    int       `json:"likes" valid:"numeric, optional"`
	Mark     Mark      `json:"mark" valid:"numeric, optional"`
}

type ReviewFull struct {
	ID     ID        `json:"id, omitempty" bson:"_id" valid:"required"`
	Title  string    `json:"title" valid:"title, stringlength(2|50)"`
	Body   string    `json:"body" valid:"description"`
	FilmID ID        `json:"film_id" valid:"numeric"`
	Author UserTrunc `json:"author, omitempty" valid:"numeric"`
	Date   time.Time `json:"date" valid:"time"`
	Likes  int       `json:"likes" valid:"numeric, optional"`
	Mark   Mark      `json:"mark" valid:"numeric, optional"`
}

func (rev *Review) AddLike() {
	rev.Likes += 1
}

func (rev *Review) RemoveLike() {
	rev.Likes -= 1
}

func (rev *Review) SetMark(mark Mark) {
	rev.Mark = mark
}
