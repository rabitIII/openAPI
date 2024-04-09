package db

import (
	"fmt"
	"testing"
)

func TestInitRedis(t *testing.T) {
	InitRedis()
	Rdb.Set("test", "value", 0)
	result, err := Rdb.Get("test").Result()
	if err != nil {
		return
	}
	fmt.Println("redis test:", result)
}
