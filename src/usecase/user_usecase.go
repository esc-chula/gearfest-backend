package usecase

import (
	"github.com/esc-chula/gearfest-backend/src/domain"
)

type UserRepository interface {
	CreateUser(*domain.User) error
	Checkin(*domain.Checkin) error
	UpdateColumn(*domain.User, string, interface{}) error
	UpdateMultipleColumns(*domain.User, map[string]interface{}) error
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

func (usecase *UserUsecase) PatchUserComplete(id string, userDTO domain.CreateUserCompletedDTO) (domain.User, error) {
	user := domain.User{
		UserID: id,
	}
	updatingMap := map[string]interface{}{
		"is_user_completed": userDTO.IsUserCompleted,
		"cocktail_id":       userDTO.CocktailID,
	}
	err := usecase.UserRepository.UpdateMultipleColumns(&user, updatingMap)
	return user, err

}

func (usecase *UserUsecase) PatchUserName(id string, userDTO domain.CreateUserNameDTO) (domain.User, error) {
	user := domain.User{
		UserID: id,
	}
	err := usecase.UserRepository.UpdateColumn(&user, "user_name", userDTO.UserName)
	return user, err

}
