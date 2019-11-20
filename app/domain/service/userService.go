package service

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}
