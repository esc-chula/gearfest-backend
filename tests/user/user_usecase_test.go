package user_tests

import (
	"errors"
	"testing"

	"github.com/esc-chula/gearfest-backend/src/domains"
	"github.com/esc-chula/gearfest-backend/src/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUsecasesTest struct {
	suite.Suite
	User          *domains.User
	CompletedUser *domains.User
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
	suite.CompletedUser = &domains.User{
		UserID:          uuid.NewString(),
		UserName:        "Test Test",
		IsUserCompleted: true,
		CocktailID:      10,
		Checkins:        []domains.Checkin{checkin},
	}
}

func (suite *UserUsecasesTest) TestGetSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	repo.On("GetById", &domains.User{}, suite.User.UserID).Return(suite.User, nil)
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
	repo.On("GetById", &domains.User{}, id).Return(nil, errors.New("User not found."))
	_, err := usecase.Get(id)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestCheckinSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	checkinDTO := domains.CreateCheckinDTO{
		LocationID: 2,
	}
	checkinInput := &domains.Checkin{
		UserID:     suite.User.UserID,
		LocationID: 2,
	}
	checkin := &domains.Checkin{
		CheckinID:  1,
		UserID:     suite.User.UserID,
		LocationID: 2,
	}
	repo.On("Checkin", checkinInput).Return(checkin, nil)
	newCheckin, err := usecase.Post(suite.User.UserID, checkinDTO)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), newCheckin, *checkin)
}

func (suite *UserUsecasesTest) TestAlreadyCheckin() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	checkinDTO := domains.CreateCheckinDTO{
		LocationID: 0,
	}
	checkinInput := &domains.Checkin{
		UserID:     suite.User.UserID,
		LocationID: 0,
	}
	repo.On("Checkin", checkinInput).Return(nil, errors.New("User already checked in."))
	_, err := usecase.Post(suite.User.UserID, checkinDTO)
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
		LocationID: 1,
	}
	checkinInput := &domains.Checkin{
		UserID:     id,
		LocationID: 1,
	}
	repo.On("Checkin", checkinInput).Return(nil, errors.New("User not found."))
	_, err := usecase.Post(id, checkinDTO)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestIsUserCompletedFalse() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	repo.On("GetField", suite.User.UserID, "is_user_completed").Return(suite.User, nil)
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
	repo.On("GetField", suite.User.UserID, "is_user_completed").Return(suite.User, nil)
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
	repo.On("GetField", id, "is_user_completed").Return(&domains.User{}, errors.New("User not found."))
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
	userInput := &domains.User{
		UserID: suite.User.UserID,
	}
	repo.On("UpdateField", userInput, suite.User.UserID, "user_name", expected.UserName).Return(&expected, nil)
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
	userInput := &domains.User{
		UserID: id,
	}
	repo.On("UpdateField", userInput, id, "user_name", "changed").Return(nil, errors.New("User not found."))
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
	updatingMap := map[string]interface{}{
		"is_user_completed": true,
		"cocktail_id":       createUserCompletedDTO.CocktailID,
	}
	repo.On("UpdateFields", &domains.User{}, suite.User.UserID, updatingMap).Return(&expected, nil)
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
	updatingMap := map[string]interface{}{
		"is_user_completed": true,
		"cocktail_id":       createUserCompletedDTO.CocktailID,
	}
	repo.On("UpdateFields", &domains.User{}, id, updatingMap).Return(nil, errors.New("User not found."))
	_, err := usecase.PatchUserComplete(id, createUserCompletedDTO)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestCreateUserSuccess() {
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
	name := "test test"
	createUserInput := &domains.User{
		UserID:   id,
		UserName: name,
		Checkins: []domains.Checkin{},
	}
	newUser := domains.User{
		UserID:          id,
		UserName:        name,
		IsUserCompleted: false,
		CocktailID:      0,
		Checkins:        []domains.Checkin{},
	}
	repo.On("CreateUser", createUserInput).Return(&newUser, nil)
	user, err := usecase.PostCreateUser(id, name)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), user, newUser)
}

func (suite *UserUsecasesTest) TestCreateUserAlreadyCreated() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	id := suite.User.UserID
	name := "test test"
	createUserInput := &domains.User{
		UserID:   id,
		UserName: name,
		Checkins: []domains.Checkin{},
	}
	repo.On("CreateUser", createUserInput).Return(nil, errors.New("User already created."))
	_, err := usecase.PostCreateUser(id, name)
	assert.NotNil(suite.T(), err)
}

func (suite *UserUsecasesTest) TestResetSuccess() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	expected := suite.CompletedUser
	expected.CocktailID = 0
	expected.IsUserCompleted = false
	updatingMap := map[string]interface{}{
		"is_user_completed": false,
		"cocktail_id":       0,
	}
	repo.On("UpdateFields", &domains.User{}, suite.CompletedUser.UserID, updatingMap).Return(expected, nil)
	user, err := usecase.ResetComplete(suite.CompletedUser.UserID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), user, *expected)
}

func (suite *UserUsecasesTest) TestResetFail() {
	repo := new(UserRepositoryMock)
	usecase := usecases.UserUsecases{
		UserRepository: repo,
	}
	var id string
	for {
		id = uuid.NewString()
		if id != suite.CompletedUser.UserID {
			break
		}
	}
	updatingMap := map[string]interface{}{
		"is_user_completed": false,
		"cocktail_id":       0,
	}
	repo.On("UpdateFields", &domains.User{}, id, updatingMap).Return(nil, errors.New("User not found."))
	_, err := usecase.ResetComplete(id)
	assert.NotNil(suite.T(), err)
}
