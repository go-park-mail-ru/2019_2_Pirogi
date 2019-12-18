package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/queryWorker"
	"github.com/labstack/echo"
)

type SearchUsecase interface {
	GetFilmsByGetParamsJSONBlob(ctx echo.Context) ([]byte, *model.Error)
	GetPersonsByGetParamsJSONBlob(ctx echo.Context) ([]byte, *model.Error)
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

func (u *searchUsecase) GetFilmsByGetParamsJSONBlob(ctx echo.Context) ([]byte, *model.Error) {
	pipeline := queryWorker.GetPipelineForMongoByContext(ctx, configs.Default.FilmTargetName)
	films, err := u.filmRepo.GetByPipeline(pipeline)
	if err != nil {
		return nil, err
	}
	var filmsFull []model.FilmFull
	for _, film := range films {
		persons := u.personRepo.GetMany(film.PersonsID)
		filmsFull = append(filmsFull, film.Full(persons))
	}
	body := modelWorker.MarshalFilmsFull(filmsFull)
	jsonBody := json.MakeJSONArray(body)
	return jsonBody, nil
}

func (u *searchUsecase) GetPersonsByGetParamsJSONBlob(ctx echo.Context) ([]byte, *model.Error) {
	pipeline := queryWorker.GetPipelineForMongoByContext(ctx, configs.Default.PersonTargetName)
	persons, err := u.personRepo.GetByPipeline(pipeline)
	if err != nil {
		return nil, err
	}
	personsTrunc := modelWorker.TruncPersons(persons)
	body := modelWorker.MarshalPersonsTrunc(personsTrunc)
	jsonBody := json.MakeJSONArray(body)
	return jsonBody, nil
}
