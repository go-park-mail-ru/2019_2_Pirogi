package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

type filmRepository struct {
	conn database.Database
}

func NewFilmRepository(conn database.Database) *filmRepository {
	return &filmRepository{
		conn: conn,
	}
}

func (r *filmRepository) Insert(newFilm model.FilmNew) *model.Error {
	return r.conn.Upsert(newFilm)
}
func (r *filmRepository) Update(film model.Film) *model.Error {
	err := r.conn.Upsert(film)
	return err
}
func (r *filmRepository) Delete(id model.ID) *model.Error {
	return r.conn.Delete(id)
}
func (r *filmRepository) Get(id model.ID) (model.Film, *model.Error) {
	film, err := r.conn.Get(id, configs.Default.FilmTargetName)
	if err != nil {
		return model.Film{}, err
	}
	return film.(model.Film), nil
}
func (r *filmRepository) GetMany(ids []model.ID) []model.Film {
	return r.conn.FindFilmsByIDs(ids)
}
func (r *filmRepository) GetByPipeline(pipeline interface{}) ([]model.Film, *model.Error) {
	filmsInterfaces, err := r.conn.GetByQuery(configs.Default.FilmsCollectionName, pipeline)
	if err != nil {
		return nil, err
	}
	var films []model.Film
	for _, filmsInterface := range filmsInterfaces {
		if film, ok := filmsInterface.(model.Film); ok {
			films = append(films, film)
		}
	}
	return films, nil
}
func (r *filmRepository) GetByTitle(title string) (model.Film, bool) {
	return r.conn.FindFilmByTitle(title)
}
