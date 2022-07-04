package store

import (
	"github.com/maryam-kermanshahani2001/send-email/model"
	"gorm.io/gorm"
)

type Student interface {
	Save(student model.Student) error
	Load(username string) (model.Student, error)
}

type SQLStudent struct {
	db *gorm.DB
}

func NewSQLStudent(db *gorm.DB) Student {
	return &SQLStudent{
		db: db,
	}
}

func (sql *SQLStudent) Save(s model.Student) error {
	result := sql.db.Create(&s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (sql *SQLStudent) Load(username string) (model.Student, error) {
	var st model.Student

	result := sql.db.Where("name = ?", username).Find(&st)
	if result.Error != nil {
		return model.Student{}, result.Error
	}

	return st, nil
}
