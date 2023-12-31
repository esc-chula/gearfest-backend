package user_tests

import (
	"errors"
	"testing"

	"github.com/esc-chula/gearfest-backend/src/domains"
	"github.com/esc-chula/gearfest-backend/src/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecasesTest struct {
	suite.Suite
	User *domains.User
}

func TestUserUsecases(t *testing.T) {
	suite.Run(t, new(UserUsecasesTest))
}

func (suite *UserUsecasesTest) SetupTest() {
	userid := uuid.NewString()
	checkin := domains.Checkin{
		CheckinID:  0,
		UserID:     userid,
		LocationID: 0,
	}
	suite.User = &domains.User{
		UserID:          userid,
		UserName:        "Test Name",
		IsUserCompleted: false,
		Checkins:        []domains.Checkin{checkin},
	}
}

func (suite *UserUsecasesTest) TestGetSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	repo.On("GetById", mock.Anything, suite.User.UserID).Return(suite.User, nil)
	user, err := usecase.Get(suite.User.UserID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), user, *suite.User)
}

func (suite *UserUsecasesTest) TestGetNotFound() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	var id string
	for {
		id = uuid.NewString()
		if id != suite.User.UserID {
			break
		}
	}
	repo.On("GetById", mock.Anything, id).Return(nil, errors.New("User not found."))
	_, err := usecase.Get(id)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestCheckinSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	checkinDTO := domains.CreateCheckinDTO{
		UserID:     suite.User.UserID,
		LocationID: 2,
	}
	checkin := &domains.Checkin{
		CheckinID:  1,
		UserID:     suite.User.UserID,
		LocationID: 2,
	}
	repo.On("Checkin", mock.Anything).Return(checkin, nil)
	newCheckin, err := usecase.Post(checkinDTO)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), newCheckin, *checkin)
}

func (suite *UserUsecasesTest) TestAlreadyCheckin() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	checkinDTO := domains.CreateCheckinDTO{
		UserID:     suite.User.UserID,
		LocationID: 0,
	}
	repo.On("Checkin", mock.Anything).Return(nil, errors.New("User already checked in."))
	_, err := usecase.Post(checkinDTO)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestCheckinUserNotFound() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	var id string
	for {
		id = uuid.NewString()
		if id != suite.User.UserID {
			break
		}
	}
	checkinDTO := domains.CreateCheckinDTO{
		UserID:     id,
		LocationID: 1,
	}
	repo.On("Checkin", mock.Anything).Return(nil, errors.New("User not found."))
	_, err := usecase.Post(checkinDTO)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestIsUserCompletedFalse() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	repo.On("GetField", suite.User.UserID, mock.Anything).Return(suite.User, nil)
	isUserCompleted, err := usecase.IsUserCompleted(suite.User.UserID)
	assert.Equal(suite.T(), isUserCompleted, false)
	assert.Nil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestIsUserCompletedTrue() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	suite.User.IsUserCompleted = true
	repo.On("GetField", suite.User.UserID, mock.Anything).Return(suite.User, nil)
	isUserCompleted, err := usecase.IsUserCompleted(suite.User.UserID)
	assert.Equal(suite.T(), isUserCompleted, true)
	assert.Nil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestIsUserCompletedNotFound() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	var id string
	for {
		id = uuid.NewString()
		if id != suite.User.UserID {
			break
		}
	}
	repo.On("GetField", id, mock.Anything).Return(&domains.User{}, errors.New("User not found."))
	isUserCompleted, err := usecase.IsUserCompleted(id)
	assert.Equal(suite.T(), isUserCompleted, false)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestPatchUserNameSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	expected := *suite.User
	expected.UserName = "changed"
	createUserNameDTO := domains.CreateUserNameDTO{
		UserName: "changed",
	}
	repo.On("UpdateField", mock.Anything, suite.User.UserID, "user_name", expected.UserName).Return(&expected, nil)
	patchedUser, err := usecase.PatchUserName(suite.User.UserID, createUserNameDTO)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), patchedUser, expected)
}

func (suite *UserUsecasesTest) TestPatchUserNameNotFound() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	var id string
	for {
		id = uuid.NewString()
		if id != suite.User.UserID {
			break
		}
	}
	createUserNameDTO := domains.CreateUserNameDTO{
		UserName: "changed",
	}
	repo.On("UpdateField", mock.Anything, id, "user_name", "changed").Return(nil, errors.New("User not found."))
	_, err := usecase.PatchUserName(id, createUserNameDTO)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestPatchUserCompleteSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	createUserCompletedDTO := domains.CreateUserCompletedDTO{
		CocktailID: 10,
	}
	expected := *suite.User
	expected.IsUserCompleted = true
	expected.CocktailID = createUserCompletedDTO.CocktailID
	repo.On("UpdateFields", mock.Anything, suite.User.UserID, mock.Anything).Return(&expected, nil)
	patchedUser, err := usecase.PatchUserComplete(suite.User.UserID, createUserCompletedDTO)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), patchedUser, expected)
}

func (suite *UserUsecasesTest) TestPatchUserCompleteNotFound() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	var id string
	for {
		id = uuid.NewString()
		if id != suite.User.UserID {
			break
		}
	}
	createUserCompletedDTO := domains.CreateUserCompletedDTO{
		CocktailID: 10,
	}
	repo.On("UpdateFields", mock.Anything, id, mock.Anything).Return(nil, errors.New("User not found."))
	_, err := usecase.PatchUserComplete(id, createUserCompletedDTO)
	assert.NotNil(suite.T(), err)
}
