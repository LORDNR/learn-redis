package product

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/lordnr/learn-redis/database"
	"github.com/lordnr/learn-redis/models"
)

var ctx = context.Background()

func Products(c *gin.Context) {

	var products []models.Product
	result := database.DB.Order("quantity desc").Limit(50).Find(&products)

	data, err := json.Marshal(products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err,
		})

		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	key := "repository::GetProducts"

	productsJson, err := rdb.Get(ctx, key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("redis")
			return
		}

	}

	err = rdb.Set(ctx, key, string(data), time.Second*10).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("database")

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"products": products,
	})
}
