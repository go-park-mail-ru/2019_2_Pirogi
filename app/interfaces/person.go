package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

type personRepository struct {
	conn database.Database
}

func NewPersonRepository(conn database.Database) *personRepository {
	return &personRepository{
		conn: conn,
	}
}

func (r *personRepository) Insert(newPerson model.PersonNew) *model.Error {
	return r.conn.Upsert(newPerson)
}
func (r *personRepository) Update(person model.Person) *model.Error {
	return r.conn.Upsert(person)
}
func (r *personRepository) Delete(id model.ID) bool {
	//TODO: работа с базой
	return true
}
func (r *personRepository) Get(id model.ID) (model.Person, *model.Error) {
	person, err := r.conn.Get(id, configs.Default.PersonTargetName)
	if err != nil {
		return model.Person{}, err
	}
	return person.(model.Person), nil
}
func (r *personRepository) GetMany(ids []model.ID) []model.Person {
	//TODO: работа с базой
	var persons []model.Person
	for _, id := range ids {
		personInterface, err := r.conn.Get(id, configs.Default.PersonTargetName)
		if err != nil {
			continue
		}
		if person, ok := personInterface.(model.Person); ok {
			persons = append(persons, person)
		}
	}
	return persons
}
func (r *personRepository) MakeTrunc(person model.Person) model.PersonTrunc {
	return person.Trunc()
}
func (r *personRepository) MakeFull(person model.Person) model.PersonFull {
	//TODO: работа с базой
	var films []model.Film
	for _, id := range person.FilmsID {
		personInterface, err := r.conn.Get(id, configs.Default.FilmTargetName)
		if err != nil {
			continue
		}
		if film, ok := personInterface.(model.Film); ok {
			films = append(films, film)
		}
	}
	return person.Full(films)
}
func (r *personRepository) GetByPipeline(pipeline interface{}) ([]model.Person, *model.Error) {
	var persons []model.Person
	personsInterface, err := r.conn.GetByQuery(configs.Default.PersonsCollectionName, pipeline)
	if err != nil {
		return nil, err
	}
	for _, personInterface := range personsInterface {
		if person, ok := personInterface.(model.Person); ok {
			persons = append(persons, person)
		}
	}
	return persons, nil
}
