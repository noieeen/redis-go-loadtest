package main

import (
	"fmt"
	"goredis/repositories"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// app := fiber.New()

	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	// time.Sleep(time.Millisecond * 10)
	// 	return c.SendString("Hello, world")
	// })

	// app.Listen(":8000")
	db := initDatabase()
	redisClient := initRedis()

	// productRepo := repositories.NewProductRepositoryDB(db)
	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	products, err := productRepo.GetProducts()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:password@tcp(localhost:3306)/infinitas")
	db, err := gorm.Open(dial, &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
