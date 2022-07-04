package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maryam-kermanshahani2001/send-email/model"
	"github.com/maryam-kermanshahani2001/send-email/store"
	"gorm.io/gorm"
	"net/http"
)

type Student struct {
	store store.Student
}

type request struct {
	Name    string  `json:"name" validate:"required"`
	Email   string  `json:"email" validate:"required,email"`
	ClassID int64   `json:"class_id" validate:"required"`
	Score   float64 `json:"score" validate:"required,gte=0,lte=20"`
}

func NewStudent(store store.Student) *Student {
	return &Student{
		store: store,
	}
}

func (s *Student) Create(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)

		return echo.ErrBadRequest
	}

	st := model.Student{
		Name:    req.Name,
		Email:   req.Email,
		ClassID: req.ClassID,
		Score:   req.Score,
	}

	if err := s.store.Save(st); err != nil {
		fmt.Println(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "")
}

func (s *Student) Find(c echo.Context) error {
	name := c.QueryParam("name")
	/*if err != nil {
		return echo.ErrBadRequest
	}*/
	println(name)

	st, err := s.store.Load(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)
}
