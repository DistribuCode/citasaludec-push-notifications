package services

import (
    "push/internal/config"
    "fmt"
)

func ProcessNotification(userID, title, body string) error {
    key := fmt.Sprintf("user:%s:notifications", userID)

    // Guarda en lista Redis
    _, err := config.RedisClient.LPush(config.Ctx, key, fmt.Sprintf("%s - %s", title, body)).Result()
    if err != nil {
        return err
    }

    // Publica en canal (pub/sub)
    _, err = config.RedisClient.Publish(config.Ctx, "notifications", fmt.Sprintf("%s|%s|%s", userID, title, body)).Result()
    return err
}
