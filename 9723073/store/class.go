package store

import (
	"github.com/maryam-kermanshahani2001/send-email/model"
	"gorm.io/gorm"
)

type Class interface {
	Save(class model.Class) error
	Load(classID int64) (model.Class, error)
}

type SQLClass struct {
	db *gorm.DB
}

func NewSQLClass(db *gorm.DB) Class {
	return &SQLClass{
		db: db,
	}
}

func (sql *SQLClass) Save(c model.Class) error {
	result := sql.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (sql *SQLClass) Load(classId int64) (model.Class, error) {
	var st model.Class

	result := sql.db.Where("class_id = ?", classId).Find(&st)
	if result.Error != nil {
		return model.Class{}, result.Error
	}

	return st, nil
}
