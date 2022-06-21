package main

// Burada todo-api programının ana fonksiyonu ve endpoint ayarlamaları yer alıyor.

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"net/http"
	_ "todo-api/api"
	"todo-api/database"
	"todo-api/handler"
)

var serverAddress = ":1323"
var e *echo.Echo

// @title Todo api
// @version 1.0
// @description Vatansoft tarafından gönderilen case için yazılmış interaktif talimatlar bütünü
// @description Todo api kullanım talimatlarını içerir.
func startServer() {
	e = echo.New()

	e.GET("/todo/:todo", handler.GetTodoUUID)
	e.GET("/todo*", handler.GetTodo)
	e.POST("/todo*", handler.CreateTodo)
	e.PUT("/todo/:todo", handler.UpdateTodo)
	e.DELETE("/todo/:todo", handler.DeleteTodo)
	e.PATCH("/todo/:todo", handler.StatusTodo)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<h1>Ana sayfadasınız</h1>
		<br><h2>Kullanım detayları için:
		<a href='https://github.com/canack/vatansoft-case'>https://github.com/canack/vatansoft-case</a></h2>`)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(serverAddress))
}

func main() {

	dbErr := database.ConnectDB()
	if dbErr != nil {
		panic(dbErr)
	}

	startServer()
}
