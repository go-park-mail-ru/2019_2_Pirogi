package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

type PersonNew struct {
	Name       string `json:"name" valid:"text, stringlength(1|50)"`
	Roles      []Role `json:"roles" valid:"roles"`
	Birthday   string `json:"birthday" valid:"date"`
	Birthplace string `json:"birthplace" valid:"text, stringlength(2|50)"`
}

func (pn *PersonNew) ToPerson(id ID) Person {
	return Person{
		ID:         id,
		Name:       pn.Name,
		Roles:      pn.Roles,
		Birthday:   pn.Birthday,
		Birthplace: pn.Birthplace,
		Genres:     []Genre{},
		FilmsID:    []ID{},
		Likes:      0,
		Images:     []Image{Image(configs.Default.DefaultImageName)},
	}
}

func (pn *PersonNew) Make(body []byte) error {
	err := pn.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(pn)
	return err
}

type Person struct {
	ID         ID      `json:"id, omitempty" bson:"_id" valid:"numeric"`
	Name       string  `json:"name" valid:"text, stringlength(1|50)"`
	Mark       Mark    `json:"mark" valid:"mark, optional"`
	Roles      []Role  `json:"type" valid:"roles"`
	Birthday   string  `json:"birthday" valid:"date"`
	Birthplace string  `json:"birthplace" valid:"text, stringlength(2|50)"`
	Genres     []Genre `json:"genres" valid:"genres, optional"`
	FilmsID    []ID    `json:"films_id" valid:"ids, optional"`
	Likes      int     `json:"likes, omitempty" valid:"numeric, optional"`
	Images     []Image `json:"images" valid:"images, optional"`
}

type PersonFull struct {
	ID         ID          `json:"id, omitempty" bson:"_id" valid:"numeric"`
	Name       string      `json:"name" valid:"text, stringlength(1|50)"`
	Mark       Mark        `json:"mark" valid:"mark, optional"`
	Roles      []Role      `json:"type" valid:"roles"`
	Birthday   string      `json:"birthday" valid:"date"`
	Birthplace string      `json:"birthplace" valid:"text, stringlength(2|50)"`
	Genres     []Genre     `json:"genres" valid:"genres, optional"`
	Films      []FilmTrunc `json:"films" valid:"optional"`
	Likes      int         `json:"likes, omitempty" valid:"numeric, optional"`
	Images     []Image     `json:"images" valid:"images, optional"`
}

type PersonTrunc struct {
	ID    ID     `json:"id, omitempty" valid:"numeric"`
	Name  string `json:"name" valid:"text, stringlength(1|50)"`
	Image Image  `json:"image" valid:"image"`
}

func (p *Person) AddLike() {
	p.Likes += 1
}

func (p *Person) RemoveLike() {
	p.Likes -= 1
}

func (p *Person) Trunc() PersonTrunc {
	return PersonTrunc{
		ID:    p.ID,
		Name:  p.Name,
		Image: p.Images[0],
	}
}

func (p *Person) Full(films []Film) PersonFull {
	var filmsTrunc []FilmTrunc
	for _, film := range films {
		filmsTrunc = append(filmsTrunc, film.Trunc())
	}
	return PersonFull{
		ID:         p.ID,
		Name:       p.Name,
		Mark:       p.Mark,
		Roles:      p.Roles,
		Birthday:   p.Birthday,
		Birthplace: p.Birthplace,
		Genres:     p.Genres,
		Films:      filmsTrunc,
		Likes:      p.Likes,
		Images:     p.Images,
	}
}

func (p *Person) HasFilmID(id ID) bool {
	for _, idx := range p.FilmsID {
		if idx == id {
			return true
		}
	}
	return false
}

func (p *Person) HasGenre(genre Genre) bool {
	for _, g := range p.Genres {
		if g == genre {
			return true
		}
	}
	return false
}
