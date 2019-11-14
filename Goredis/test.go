package main



import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)



var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}
func main(){
	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}
