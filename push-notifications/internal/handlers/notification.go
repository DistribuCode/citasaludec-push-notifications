package handlers

import (
    "encoding/json"
    "net/http"
    "push/internal/services"
)

type NotificationRequest struct {
    UserID string `json:"user_id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
    var req NotificationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    // Llama al servicio para guardar y publicar
    if err := services.ProcessNotification(req.UserID, req.Title, req.Body); err != nil {
        http.Error(w, "Error procesando notificación", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Notificación enviada"})
}
