package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type FilmRepository interface {
	Insert(newFilm model.FilmNew) *model.Error
	Update(film model.Film) *model.Error
	Delete(id model.ID) *model.Error
	Get(id model.ID) (model.Film, *model.Error)
	GetMany(ids []model.ID) []model.Film
	GetByPipeline(pipeline interface{}) ([]model.Film, *model.Error)
	GetByTitle(title string) (model.Film, bool)
	GetRelated(filmFull model.FilmFull) ([]model.FilmTrunc, *model.Error)
}
