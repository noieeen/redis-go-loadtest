package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv  services.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) catalogHandlerRedis {
	return catalogHandlerRedis{catalogSrv, redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {

	// Get redis
	key := "handler::GetProducts"
	if responseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("Redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	// Service
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Redis Set
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("Database")
	return c.JSON(response)
}
