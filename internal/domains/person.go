package domains

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/security"
	"golang.org/x/net/html"
)

type PersonRepository interface {
	Insert(newPerson PersonNew) (ID, error)
	Update(id ID, person Person) error
	Delete(id ID) bool
	Get(id ID) Person
	GetMany(target Target, id ID) []Person
	MakeTrunc(person Person) PersonTrunc
	MakeFull(person Person) PersonFull
}

type PersonNew struct {
	Name       string `json:"name" valid:"text, stringlength(1|50)"`
	Roles      []Role `json:"roles" valid:"roles"`
	Birthday   string `json:"birthday" valid:"date"`
	Birthplace string `json:"birthplace" valid:"text, stringlength(2|50)"`
}

func (np *PersonNew) ToPerson(id ID) Person {
	return Person{
		ID:         id,
		Name:       html.EscapeString(np.Name),
		Roles:      security.XSSFilterRoles(np.Roles),
		Birthday:   html.EscapeString(np.Birthday),
		Birthplace: html.EscapeString(np.Birthplace),
		Genres:     []Genre{},
		FilmsID:    []ID{},
		Likes:      0,
		Images:     []Image{configs.Default.DefaultImageName},
	}
}

func (np *PersonNew) Make(body []byte) error {
	err := np.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(np)
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
