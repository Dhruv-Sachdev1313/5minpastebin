package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	Health() map[string]string
	SetKeyWithExpiration(string, string) error
	GetData(string) (string, error)
}

type service struct {
	redis *redis.Client
}

var (
	host = os.Getenv("REDIS_DB_HOST")
	port = os.Getenv("REDIS_DB_PORT")
)

func New() Service {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       0,
	})

	return &service{
		redis: rdb,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.redis.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("redis down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) SetKeyWithExpiration(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	status := s.redis.SetEx(ctx, key, value, time.Minute*5)
	return status.Err()
}

func (s *service) GetData(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cmd := s.redis.Get(ctx, key)
	return cmd.Result()
}
