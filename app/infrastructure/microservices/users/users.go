package users

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/users/protobuf"
)

func NewUsersManager(userRepo repository.UserRepository, cookieRepo repository.CookieRepository) *usersManager {
	return &usersManager{
		userRepo:   userRepo,
		cookieRepo: cookieRepo,
	}
}

type usersManager struct {
	userRepo   repository.UserRepository
	cookieRepo repository.CookieRepository
}

func (u *usersManager) GetByCookieValue(ctx context.Context, req *v1.CookieValue) (*v1.User, error) {
	user, err := u.cookieRepo.GetUserByCookieValue(req.CookieValue)
	if err != nil {
		return &v1.User{}, err.Common()
	}
	protoUser := user.ToProtobuf()
	return &protoUser, nil
}

func (u *usersManager) GetByID(ctx context.Context, req *v1.ID) (*v1.User, error) {
	user, err := u.userRepo.Get(model.ID(req.ID))
	if err != nil {
		return &v1.User{}, err.Common()
	}
	protoUser := user.ToProtobuf()
	return &protoUser, nil
}

func (u *usersManager) Create(ctx context.Context, user *v1.UserNew) (*v1.Nothing, error) {
	var userNewModel model.UserNew
	userNewModel.FromProtobuf(*user)
	err := u.userRepo.Insert(userNewModel)
	return &v1.Nothing{}, err.Common()
}

func (u *usersManager) Update(req context.Context, user *v1.User) (*v1.Nothing, error) {
	var userModel model.User
	userModel.FromProtobuf(*user)
	err := u.userRepo.Update(userModel)
	return &v1.Nothing{}, err.Common()
}
