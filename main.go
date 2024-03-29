package main

import (
	"Projects/test_exercise/config"
	"Projects/test_exercise/handler"

	"github.com/gin-gonic/gin"
)

/*

Необходимо создать Http сервис - key-value хранилище.
Сервис должен содержать четыре метода в апи:
- Upsert (вставка либо обновление)
- Delete
- Get
- List
Хранить данные можно просто в оперативной памяти при помощи map.
C:\Users\Anton\go\src\Projects\test_exercise\Dockerfile
*/

func main() {

	r := gin.Default()

	handler := handler.NewHandler()
	config := config.NewConfig("app/config.json")

	r.GET("/get/:key", handler.Get)
	r.GET("/list", handler.List)
	r.DELETE("/delete/:key", handler.Delete)
	r.POST("/upsert", handler.Upsert)
	r.GET("/")
	r.GET("/favicon.ico")
	r.GET("/index.html")
	r.GET("/redis", handler.Redis)
	r.GET("/redis-start", handler.RedisStart)
	r.Run(":" + config.Port)
	//
}
