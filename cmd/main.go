package main

import (
    "log"
    "net/http"
    "push/internal/config"
    "push/internal/handlers"
    "push/internal/middleware"

    "github.com/gorilla/mux"
)

func main() {
    // Inicializa conexión Redis
    config.InitRedis()

    // Router
    r := mux.NewRouter()

    // Middlewares
    r.Use(middleware.CorsMiddleware)
    r.Use(middleware.JwtMiddleware)

    // Rutas
    r.HandleFunc("/notifications", handlers.CreateNotification).Methods("POST")

    // Levanta servidor
    log.Println("🚀 Push Notification Service corriendo en :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
