package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maryam-kermanshahani2001/send-email/model"
	"github.com/maryam-kermanshahani2001/send-email/store"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type classRequest struct {
	ClassId    int64  `json:"class_id" validate:"required"`
	ClassName  string `json:"name" validate:"required"`
	Instructor string `json:"inst_name" validate:"required"`
}

type Class struct {
	store store.Class
}

func NewClass(store store.Class) *Class {
	return &Class{
		store: store,
	}
}

func (cls *Class) CreateClass(c echo.Context) error {
	var req classRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)

		return echo.ErrBadRequest
	}

	cl := model.Class{
		ClassId:    req.ClassId,
		ClassName:  req.ClassName,
		Instructor: req.Instructor,
	}

	if err := cls.store.Save(cl); err != nil {
		fmt.Println(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, "")
}

func (cls *Class) Find(c echo.Context) error {
	classId, err := strconv.Atoi(c.QueryParam("class_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	println(classId)

	st, err := cls.store.Load(int64(classId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)
}
