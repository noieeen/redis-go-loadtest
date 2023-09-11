package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
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
	_ = db
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	// productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	// productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogService(productRepo)
	// productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	// productHandler := handlers.NewCatalogHandler(productService)
	productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient)

	// products, err := productService.GetProducts()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(products)

	// fiber by Handler
	app := fiber.New()

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":8000")
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
