package main

import (
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Student{Name: "Neeraj Yadav", Age: 26})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("7541", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("7541").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
