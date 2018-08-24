package main

import (
	"github.com/ShingoYadomoto/go-echo-vue-single-page-app/handlers"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks)

	e.Start(":8080")
}
