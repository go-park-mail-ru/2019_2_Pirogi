package models

type NewPerson struct {
	Name       string      `json:"name"`
	Roles      []Role      `json:"type"`
	Birthday   string      `json:"birthday"`
	Birthplace string      `json:"birthplace"`
	Genres     []Genre     `json:"genres_id"`
	Films      []FilmTrunc `json:"films_id"`
	Likes      int         `json:"rating, omitempty"`
	Images     []Image     `json:"images_id"`
}

type Person struct {
	PersonTrunc `valid:"required"`
	Roles       []Role      `json:"type" valid:"roles"`
	Birthday    string      `json:"birthday" valid:"date"`
	Birthplace  string      `json:"birthplace" valid:"alphanum, stringlength(2|50)"`
	Genres      []Genre     `json:"genres" valid:"genres"`
	Films       []FilmTrunc `json:"films" valid:"films_trunc"`
	Likes       int         `json:"likes, omitempty" valid:"numeric, optional"`
	Images      []Image     `json:"images" valid:"images, optional"`
}

type PersonTrunc struct {
	ID   ID     `json:"id, omitempty" valid:"numeric"`
	Name string `json:"name" valid:"alpha, stringlength(1|50)"`
	Mark Mark   `json:"mark" valid:"mark, optional"`
}
