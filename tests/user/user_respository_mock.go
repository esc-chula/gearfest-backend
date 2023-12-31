package user_tests

import (
	"github.com/esc-chula/gearfest-backend/src/domains"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user *domains.User) error {
	args := m.Called(user)
	if args.Get(0) != nil {
		*user = *args.Get(0).(*domains.User)
	}
	return args.Error(1)
}

func (m *UserRepositoryMock) Checkin(checkin *domains.Checkin) error {
	args := m.Called(checkin)
	if args.Get(0) != nil {
		*checkin = *args.Get(0).(*domains.Checkin)
	}
	return args.Error(1)
}

func (m *UserRepositoryMock) GetField(id string, field string) (*domains.User, error) {
	args := m.Called(id, field)
	return args.Get(0).(*domains.User), args.Error(1)
}

func (m *UserRepositoryMock) UpdateField(user *domains.User, id string, field string, value interface{}) error {
	args := m.Called(user, id, field, value)
	if args.Get(0) != nil {
		*user = *args.Get(0).(*domains.User)
	}
	return args.Error(1)
}

func (m *UserRepositoryMock) UpdateFields(user *domains.User, id string, fields map[string]interface{}) error {
	args := m.Called(user, id, fields)
	if args.Get(0) != nil {
		*user = *args.Get(0).(*domains.User)
	}
	return args.Error(1)
}

func (m *UserRepositoryMock) GetById(user *domains.User, id string) error {
	args := m.Called(user, id)
	if args.Get(0) != nil {
		*user = *args.Get(0).(*domains.User)
	}
	return args.Error(1)
}
