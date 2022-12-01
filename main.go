package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lordnr/learn-redis/config"
	"github.com/lordnr/learn-redis/database"
	"github.com/lordnr/learn-redis/routes"
)

// var ctx = context.Background()

func init() {
	config.LoadEnvironment()
	database.ConnectDB()

}

// func initRedis() {
// 	// rdb := redis.NewClient(&redis.Options{
// 	// 	Addr: "localhost:6379",
// 	// })
// 	// err := rdb.Set(ctx, "key", "value", 0).Err()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		time.Sleep(time.Millisecond * 10)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	routes.ProductRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
