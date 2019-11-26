package users

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/users/protobuf"
	"go.uber.org/zap"
)

func NewUsersManager(userRepo repository.UserRepository, cookieRepo repository.CookieRepository) *usersManager {
	zap.S().Debug("Creating new users manager")
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
	zap.S().Debug("creating new user with email: ", user.Email)
	userNewModel.FromProtobuf(*user)
	err := u.userRepo.Insert(userNewModel)
	zap.S().Debug(err)
	return &v1.Nothing{}, err.Common()
}

func (u *usersManager) Update(ctx context.Context, user *v1.User) (*v1.Nothing, error) {
	var userModel model.User
	userModel.FromProtobuf(*user)
	err := u.userRepo.Update(userModel)
	return &v1.Nothing{}, err.Common()
}

func (u *usersManager) GetByEmail(ctx context.Context, email *v1.Email) (*v1.User, error) {
	user, err := u.userRepo.GetByEmail(email.Email)
	if err != nil {
		return &v1.User{}, err.Common()
	}
	protoUser := user.ToProtobuf()
	return &protoUser, nil
}
