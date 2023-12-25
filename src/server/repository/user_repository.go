package repository

import (
	"github.com/esc-chula/gearfest-backend/src/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create user in database
func (repo *UserRepository) CreateUser(user *domain.User) error {
	result := repo.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Create checkin
func (repo *UserRepository) Checkin(checkin *domain.Checkin) error {
	result := repo.db.Create(checkin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update a column of user in database with column_name and value
func (repo *UserRepository) UpdateColumn(user *domain.User, column_name string, value interface{}) error {
	result := repo.db.Model(user).Clauses(clause.Returning{}).Update(column_name, value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update some columns of user in database
func (repo *UserRepository) UpdateMultipleColumns(user *domain.User, columns map[string]interface{}) error {
	result := repo.db.Model(user).Clauses(clause.Returning{}).Updates(columns)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get user by id
func (repo *UserRepository) GetById(user *domain.User, id string) error {
	result := repo.db.First(user, "user_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get user with checkins
func (repo *UserRepository) GetWithCheckins(user *domain.User, id string) error {
	result := repo.db.Model(&domain.User{}).Preload("Checkins").First(user, "user_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
