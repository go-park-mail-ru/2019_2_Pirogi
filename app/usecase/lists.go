package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

type ListsUsecase interface {
	CreateList(list model.ListNew) *model.Error
	AddFilmToList(id model.ID, filmID model.ID) *model.Error
	GetListsByUserID(userID model.ID) ([]model.ID, *model.Error)
}

type listsUsecase struct {
	cookieRepo repository.CookieRepository
	listsRepo  repository.ListRepository
}

func (l listsUsecase) CreateListByContext(ctx echo.Context) *model.Error {
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

	list.UserID = user.ID
	err = l.listsRepo.Insert(list)
	if err != nil {
		return err
	}
	return err
}

func (l listsUsecase) AddFilmToUserList(ctx echo.Context) *model.Error {
	user, err := l.cookieRepo.GetUserByContext(ctx)
	if err != nil {
		return err
	}
	body, err := network.ReadBody(ctx)
	if err != nil {
		return err
	}
	var
	list, err :=
}

func (l listsUsecase) GetListsByUserID(userID model.ID) ([]model.ID, *model.Error) {
	panic("implement me")
}

func NewListsUsecase(cookieRepo repository.CookieRepository, listsRepo repository.ListRepository) *listsUsecase {
	return &listsUsecase{
		cookieRepo: cookieRepo,
		listsRepo:  listsRepo,
	}
}
