package model

type Class struct {
	ClassId    int64  `json:"class_id" validate:"required"`
	ClassName  string `json:"name" validate:"required"`
	Instructor string `json:"inst_name" validate:"required"`
}
