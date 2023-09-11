package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ProductRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepositoryRedis {
	db.AutoMigrate()
	mockData(db)
	return ProductRepositoryRedis{db, redisClient}
}

func (r ProductRepositoryRedis) GetProducts() (products []product, err error) {
	// Redis Key
	key := "repository::GetProducts"

	// Redis GetProducts
	productJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productJson), &products)
		if err == nil {
			fmt.Println("Redis")
			return products, nil
		}
	}

	// Database
	err = r.db.Order("quantity desc").Limit(50).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// Redis Set
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database")

	return products, nil
}
