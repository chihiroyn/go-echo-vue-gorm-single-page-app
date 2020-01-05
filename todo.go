package main

import (
	"github.com/ShingoYadomoto/go-echo-vue-single-page-app/handlers"
	"github.com/labstack/echo"
	"go-echo-vue-gorm-single-page-app/models"
)

func main() {
	models.InitDB()

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks)

	e.Start(":8080")
}
