package domains

type PersonRepository interface {
	Insert(newPerson NewPerson) (ID, error)
	Update(id ID, person Person) error
	Delete(id ID) bool
	Get(id ID) Person
	GetMany(target Target, id ID) []Person
	MakeTrunc(person Person) PersonTrunc
	MakeFull(person Person) PersonFull
}

type NewPerson struct {
	Name       string `json:"name" valid:"text, stringlength(1|50)"`
	Roles      []Role `json:"roles" valid:"roles"`
	Birthday   string `json:"birthday" valid:"date"`
	Birthplace string `json:"birthplace" valid:"text, stringlength(2|50)"`
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
