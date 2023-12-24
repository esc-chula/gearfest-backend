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
	//pass user id so when send to UserRepository it'll find user with the id
	user := domain.User{
		UserID: id,
	}
	//send the value by address to assign all other field to user that we have created
	err := usecase.UserRepository.GetByPrimaryKey(&user)
	return user, err
}

func (usecase *UserUsecase) Post(CheckinDTO domain.CreateCheckinDTO) (domain.Checkin,error) {

	checkin := domain.Checkin{
		UserID: CheckinDTO.UserID,
		LocationID: CheckinDTO.LocationID,
	
	}

	err := usecase.UserRepository.Create(&checkin)
	return checkin,err
}

