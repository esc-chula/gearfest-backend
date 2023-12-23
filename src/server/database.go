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

func (handler *SqlHandler) Create(obj interface{}) error {
	result := handler.db.Session(&gorm.Session{}).Create(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (handler *SqlHandler) UpdateColumn(obj interface{}, column_name string, value interface{}) error {
	result := handler.db.Session(&gorm.Session{}).Model(&obj).Update(column_name, value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (handler *SqlHandler) GetByPrimaryKey(obj interface{}) error {
	result := handler.db.Session(&gorm.Session{}).First(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (handler *SqlHandler) GetWithAssociations(obj interface{}) error {
	result := handler.db.Session(&gorm.Session{}).Preload(clause.Associations).First(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
