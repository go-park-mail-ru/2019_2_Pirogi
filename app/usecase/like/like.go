package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
)

type LikeUsecase interface {
	Set(ctx echo.Context) *model.Error
	Unset(id model.ID) *model.Error
	List(ids []model.ID) []model.Person
}

type likeUsecase struct {
	likeRepo repository.LikeRepository
	filmRepo repository.FilmRepository
	userRepo repository.UserRepository
}

func NewLikeUsecase(likeRepo repository.LikeRepository, userRepo repository.UserRepository, filmRepo repository.FilmRepository) *likeUsecase {
	return &likeUsecase{
		likeRepo: likeRepo,
		filmRepo: filmRepo,
		userRepo: userRepo,
	}
}

func (u *likeUsecase) Set(ctx echo.Context) *model.Error {
	user, ok := u.userRepo.GetByContext(ctx)
	if !ok {
		return model.NewError(400, "not found")
	}
	body, err := network.ReadBody(ctx)
	if err != nil {
		return err
	}
	like := model.Like{}
	e := like.UnmarshalJSON(body)
	if e != nil {
		return model.NewError(400, e.Error())
	}
	like.UserID = user.ID
	return u.likeRepo.Insert(like)
}
