package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maryam-kermanshahani2001/send-email/handler"
	"github.com/maryam-kermanshahani2001/send-email/model"
	"github.com/maryam-kermanshahani2001/send-email/store"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()

	stdb, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	cldb, err := gorm.Open(sqlite.Open("classes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// disabling gorm logs
	// db.Logger.LogMode(logger.Silent)

	if err := stdb.Migrator().AutoMigrate(&model.Student{}); err != nil {
		log.Fatal(err)
	}

	if err := cldb.Migrator().AutoMigrate(&model.Class{}); err != nil {
		log.Fatal(err)
	}
	//e.GET("/hello", handler.Hello)

	s := handler.NewStudent(store.NewSQLStudent(stdb))
	c := handler.NewClass(store.NewSQLClass(cldb))
	e.POST("/class", c.CreateClass)
	e.GET("/class", c.Find)
	e.POST("/student", s.Create)
	e.GET("/student", s.Find)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}

}
