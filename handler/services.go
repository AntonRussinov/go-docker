package handler

import (
	"Projects/test_exercise/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

//map для хранения данных
var m = make(map[string]string)

//Get для получения данных по ключу.
func (h *Handler) Get(c *gin.Context) {
	key := c.Param("key")
	val := m[key]
	resp := model.Response{Key: key, Value: val}
	c.JSON(http.StatusOK, resp)
}

//Delete для удаления данных по ключу.
func (h *Handler) Delete(c *gin.Context) {
	key := c.Param("key")
	delete(m, key)
	c.Status(http.StatusNoContent)
}

//Upsert для добавления или обновления записи.
func (h *Handler) Upsert(c *gin.Context) {
	insert := model.Response{}
	err := c.ShouldBindJSON(&insert)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error(), HTTPStatus: http.StatusBadRequest})
	}
	m[insert.Key] = insert.Value
	//resp := model.Response{Key: key, Value: val}
	c.Status(http.StatusNoContent)
}

//List получение всех записей из map
func (h *Handler) List(c *gin.Context) {
	resp := []model.Response{}
	for key, val := range m {
		resp = append(resp, model.Response{Key: key, Value: val})
	}
	c.JSON(http.StatusOK, resp)
}

//Redis
func (h *Handler) Redis(c *gin.Context) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get("visit").Result()
	if err != nil {
		err = rdb.Set("visit", 1, 0).Err()
		if err != nil {
			c.JSON(http.StatusBadGateway, err)
		}
		c.JSON(http.StatusOK, 1)
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
	}
	err = rdb.Set("visit", valInt+1, 0).Err()
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
	}
	//fmt.Println(val)
	c.JSON(http.StatusOK, val)

}
