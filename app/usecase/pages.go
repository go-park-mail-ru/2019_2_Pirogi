package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/queryWorker"
)

type PagesUsecase interface {
	GetIndexPageJSONBlob() ([]byte, *model.Error)
}

func NewPagesUsecase(filmRepo repository.FilmRepository, personRepo repository.PersonRepository) *pagesUsecase {
	return &pagesUsecase{
		filmRepo:   filmRepo,
		personRepo: personRepo,
	}
}

type pagesUsecase struct {
	filmRepo   repository.FilmRepository
	personRepo repository.PersonRepository
}

func (u *pagesUsecase) GetIndexPageJSONBlob() ([]byte, *model.Error) {
	//TODO: добавить бы логики тута
	pipeline := queryWorker.GetCustomPipelineForMongo(configs.Default.DefaultEntriesLimit+5,
		0, configs.Default.FilmTargetName)
	filmsNew, err := u.filmRepo.GetByPipeline(pipeline)
	if err != nil {
		return nil, err
	}
	pipeline = queryWorker.GetCustomPipelineForMongo(configs.Default.DefaultEntriesLimit+20,
		0, configs.Default.FilmTargetName)
	filmsForUser, err := u.filmRepo.GetByPipeline(pipeline)
	if err != nil {
		return nil, err
	}
	trailers := modelWorker.MakeTrailersList(filmsNew)
	trailersBody := modelWorker.MarshalTrailers(trailers)

	filmsNewTrunc := modelWorker.TruncFilms(filmsNew)
	filmsNewBody := modelWorker.MarshalFilmsTrunc(filmsNewTrunc)

	filmsForUserTrunc := modelWorker.TruncFilms(filmsForUser[15:])
	filmsForUserBody := modelWorker.MarshalFilmsTrunc(filmsForUserTrunc)

	return json.UnionToJSON([]string{"filmsNew", "filmsForUser", "trailers"}, filmsNewBody, filmsForUserBody, trailersBody), nil
}
