package controllers

import (
    "encoding/json"
    "net/http"

    "yourapp/models"
    "yourapp/services"
)

// UserController definisce le funzioni per gestire le richieste relative agli utenti
type UserController struct {
    UserService *services.UserService
}

// NewUserController crea una nuova istanza di UserController
func NewUserController(userService *services.UserService) *UserController {
    return &UserController{UserService: userService}
}

// GetUserHandler gestisce le richieste GET /users
func (uc *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    users, err := uc.UserService.GetUsers()
    if err != nil {
        http.Error(w, "Errore durante il recupero degli utenti", http.StatusInternalServerError)
        return
    }

    // Converti gli utenti in JSON e invia la risposta
    json.NewEncoder(w).Encode(users)
}
