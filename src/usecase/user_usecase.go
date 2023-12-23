package usecase

import (
	"github.com/esc-chula/gearfest-backend/src/domain"
	"github.com/esc-chula/gearfest-backend/src/interfaces"
)
type UserRepository struct {
	interfaces.SqlHandler
}

type UserUsecase struct {
	UserRepository UserRepository
}

func (usecase *UserUsecase) Get(id string) (domain.User, error) {
	user := domain.User{
		UserID: id,
	}
	err := usecase.UserRepository.GetByPrimaryKey(&user)
	return user, err
}