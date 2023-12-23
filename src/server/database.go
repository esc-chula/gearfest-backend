package server

import (
	"github.com/esc-chula/gearfest-backend/src/interfaces"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler(db *gorm.DB) interfaces.SqlHandler {
	SqlHandler := new(SqlHandler)
	SqlHandler.db = db
	return SqlHandler
}

//Create obj in database
func (handler *SqlHandler) Create(obj interface{}) error {
	result := handler.db.Session(&gorm.Session{}).Create(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Update column of obj in database with column_name and value
func (handler *SqlHandler) UpdateColumn(obj interface{}, column_name string, value interface{}) error {
	result := handler.db.Session(&gorm.Session{}).Model(&obj).Update(column_name, value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Get data of object by primary key
func (handler *SqlHandler) GetByPrimaryKey(obj interface{}) error {
	result := handler.db.Session(&gorm.Session{}).First(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Get data of object with associations
func (handler *SqlHandler) GetWithAssociations(obj interface{}) error {
	result := handler.db.Session(&gorm.Session{}).Preload(clause.Associations).First(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
