package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

// rdb := redis.NewClient(&redis.Options{
// 	Addr:     "localhost:6379",
// 	Password: "", // no password set
// 	DB:       0,  // use default DB
// })

func ConnectRedis() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val)

}
