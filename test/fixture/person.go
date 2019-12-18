package fixture

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

var PersonID = model.ID(1)
var Name = "test person"
var Roles = []model.Role{"actor", "producer"}
var Birthday = "09.12.1998"
var Birthplace = "USA"
var FilmsID = []model.ID{1, 2, 3}
var Images = []model.Image{Image}

var Persons = []model.Person{Person, Person}

var Person = model.Person{
	ID:         PersonID,
	Name:       Name,
	Mark:       Mark,
	Roles:      Roles,
	Birthday:   Birthday,
	Birthplace: Birthplace,
	Genres:     Genres,
	FilmsID:    FilmsID,
	Images:     Images,
}

var PersonTrunc = Person.Trunc()
var PersonFull = Person.Full(Films)
