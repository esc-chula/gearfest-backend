package repositories

import (
	"github.com/esc-chula/gearfest-backend/src/domains"
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
func (repo *UserRepository) CreateUser(user *domains.User) error {
	result := repo.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Create checkin
func (repo *UserRepository) Checkin(checkin *domains.Checkin) error {
	result := repo.db.Create(checkin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update a column of user in database with column_name and value
func (repo *UserRepository) UpdateField(user *domains.User, id string, column_name string, value interface{}) error {
	result := repo.db.Model(&domains.User{}).Where("user_id = ?", id).Update(column_name, value)
	if result.Error != nil {
		return result.Error
	}
	return repo.GetById(user, id)
}

// Update some columns of user in database
func (repo *UserRepository) UpdateFields(user *domains.User, id string, columns map[string]interface{}) error {
	result := repo.db.Model(&domains.User{}).Where("user_id = ? ", id).Updates(columns)
	if result.Error != nil {
		return result.Error
	}
	return repo.GetById(user, id)
}

// Get user by id
func (repo *UserRepository) GetById(user *domains.User, id string) error {
	result := repo.db.Model(&domains.User{}).Preload(clause.Associations).First(user, "user_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
