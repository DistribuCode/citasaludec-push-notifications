package middleware

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v4"
    "os"
)

func JwtMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
        if tokenString == "" {
            http.Error(w, "Token requerido", http.StatusUnauthorized)
            return
        }

        _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil {
            http.Error(w, "Token inválido", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
