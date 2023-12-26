package usecases

import "github.com/esc-chula/gearfest-backend/src/domains"

type UserRepository interface {
	CreateUser(*domains.User) error
	Checkin(*domains.Checkin) error
	UpdateField(*domains.User, string, string, interface{}) error
	UpdateFields(*domains.User, string, map[string]interface{}) error
	GetById(*domains.User, string) error
}

type UserUsecases struct {
	UserRepository UserRepository
}

func (usecase *UserUsecases) Get(id string) (domains.User, error) {
	user := domains.User{}
	err := usecase.UserRepository.GetById(&user, id)
	return user, err
}

func (usecase *UserUsecases) Post(CheckinDTO domains.CreateCheckinDTO) (domains.Checkin, error) {

	checkin := domains.Checkin{
		UserID:     CheckinDTO.UserID,
		LocationID: CheckinDTO.LocationID,
	}

	err := usecase.UserRepository.Checkin(&checkin)
	return checkin, err
}

func (usecase *UserUsecases) PatchUserComplete(id string, userDTO domains.CreateUserCompletedDTO) (domains.User, error) {
	user := domains.User{}
	updatingMap := map[string]interface{}{
		"is_user_completed": userDTO.IsUserCompleted,
		"cocktail_id":       userDTO.CocktailID,
	}
	err := usecase.UserRepository.UpdateFields(&user, id, updatingMap)
	return user, err

}

func (usecase *UserUsecases) PatchUserName(id string, userDTO domains.CreateUserNameDTO) (domains.User, error) {
	user := domains.User{
		UserID: id,
	}
	err := usecase.UserRepository.UpdateField(&user, id, "user_name", userDTO.UserName)
	return user, err

}
