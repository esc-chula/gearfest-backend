package usecase

import (
	"github.com/esc-chula/gearfest-backend/src/domain"
)

type UserRepository interface {
	CreateUser(*domain.User) error
	Checkin(*domain.Checkin) error
	UpdateColumn(*domain.User, string, interface{}) error
	GetById(*domain.User, string) error
	GetWithCheckins(*domain.User, string) error
}

type UserUsecase struct {
	UserRepository UserRepository
}

func (usecase *UserUsecase) Get(id string) (domain.User, error) {
	user := domain.User{}
	//send the value by address to assign all other field to user that we have created
	err := usecase.UserRepository.GetById(&user, id)
	return user, err
}

func (usecase *UserUsecase) Post(CheckinDTO domain.CreateCheckinDTO) (domain.Checkin, error) {

	checkin := domain.Checkin{
		UserID:     CheckinDTO.UserID,
		LocationID: CheckinDTO.LocationID,
	}

	err := usecase.UserRepository.Checkin(&checkin)
	return checkin, err
}
