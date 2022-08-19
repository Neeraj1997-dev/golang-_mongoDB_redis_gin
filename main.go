package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"backend/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	server      *gin.Engine
	ctx         context.Context
	redisclient *redis.Client
)

func init() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	ctx = context.TODO()
	// Connect to Redis
	redisclient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})

	if _, err := redisclient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err = redisclient.Set(ctx, "test", "Welcome to Golang with Redis", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")

	//Create the Gin Engine instance
	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}
	value, err := redisclient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	log.Fatal(server.Run(":" + config.Port))
}
