package model

type Student struct {
	Name    string  `json:"name" validate:"required"`
	Email   string  `json:"email" validate:"required,email"`
	ClassID int64   `json:"class_id" validate:"required"`
	Score   float64 `json:"score" validate:"required,gte=0,lte=20"`
}
