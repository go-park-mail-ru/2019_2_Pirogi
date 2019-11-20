package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/queryWorker"
	"github.com/labstack/echo"
)

type SearchUsecase interface {
	GetFilmsByGetParams(ctx echo.Context) ([]model.Film, model.Error)
	GetPersonsByGetParams(ctx echo.Context) ([]model.Person, model.Error)
}

func NewSearchUsecase(filmRepo repository.FilmRepository, personRepo repository.PersonRepository) *searchUsecase {
	return &searchUsecase{
		filmRepo:   filmRepo,
		personRepo: personRepo,
	}
}

type searchUsecase struct {
	filmRepo   repository.FilmRepository
	personRepo repository.PersonRepository
}

func (u *searchUsecase) GetFilmsByGetParams(ctx echo.Context) ([]model.Film, *model.Error) {
	pipeline := queryWorker.GetPipelineForMongo(ctx, "films")
	return u.filmRepo.GetByPipeline(pipeline)
}

func (u *searchUsecase) GetPersonsByGetParams(ctx echo.Context) ([]model.Person, *model.Error) {
	pipeline := queryWorker.GetPipelineForMongo(ctx, "films")
	return u.personRepo.GetByPipeline(pipeline)
}
