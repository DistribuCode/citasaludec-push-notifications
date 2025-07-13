package config

import (
    "github.com/go-redis/redis/v8"
    "context"
    "log"
    "os"
)

var (
    RedisClient *redis.Client
    Ctx = context.Background()
)

func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: os.Getenv("REDIS_ADDR"),
        Password: "",
        DB: 0,
    })

    _, err := RedisClient.Ping(Ctx).Result()
    if err != nil {
        log.Fatalf("No se pudo conectar a Redis: %v", err)
    }

    log.Println("✅ Conectado a Redis")
}
