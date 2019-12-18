package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type ListsUsecase interface {
	CreateOrUpdateList(ctx echo.Context) *model.Error
	GetListsByUserCtx(ctx echo.Context) ([]model.ListFull, *model.Error)
}

type listsUsecase struct {
	cookieRepo repository.CookieRepository
	listsRepo  repository.ListRepository
	filmsRepo  repository.FilmRepository
}

func (l listsUsecase) CreateOrUpdateList(ctx echo.Context) *model.Error {
	user, err := l.cookieRepo.GetUserByContext(ctx)
	if err != nil {
		return err
	}
	body, err := network.ReadBody(ctx)
	if err != nil {
		return err
	}
	var list model.ListNew
	e := list.UnmarshalJSON(body)
	if e != nil {
		return model.NewError(400, e.Error())
	}
	listOld, err := l.listsRepo.GetByUserIDAndTitle(user.ID, list.Title)
	if err != nil {
		list.UserID = user.ID
		return l.listsRepo.Insert(list)
	}
	listOld.FilmsID = append(listOld.FilmsID, list.FilmID)
	return l.listsRepo.Update(listOld)
}

func (l listsUsecase) GetListsByUserCtx(ctx echo.Context) ([]model.ListFull, *model.Error) {
	user, err := l.cookieRepo.GetUserByContext(ctx)
	if err != nil {
		return nil, err
	}
	lists, err := l.listsRepo.GetByUserID(user.ID)
	zap.S().Debug(lists)
	if err != nil {
		return nil, err
	}
	var listsFull []model.ListFull
	for _, list := range lists {
		var listFull model.ListFull
		listFull.Title = list.Title
		listFull.ID = list.ID
		listFull.UserID = list.UserID
		films := l.filmsRepo.GetMany(list.FilmsID)
		filmsTrunc := modelWorker.TruncFilms(films)
		listFull.Films = filmsTrunc
		listsFull = append(listsFull, listFull)
	}
	zap.S().Debug(listsFull)
	return listsFull, nil
}

func NewListsUsecase(cookieRepo repository.CookieRepository, listsRepo repository.ListRepository,
	filmsRepo repository.FilmRepository) *listsUsecase {
	return &listsUsecase{
		cookieRepo: cookieRepo,
		listsRepo:  listsRepo,
		filmsRepo:  filmsRepo,
	}
}
