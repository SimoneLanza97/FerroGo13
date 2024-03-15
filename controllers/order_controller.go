package controllers

import (
    "encoding/json"
    "net/http"

    "yourapp/models"
    "yourapp/services"
)

// OrderController definisce le funzioni per gestire le richieste relative agli ordini
type OrderController struct {
    OrderService *services.OrderService
}

// NewOrderController crea una nuova istanza di OrderController
func NewOrderController(orderService *services.OrderService) *OrderController {
    return &OrderController{OrderService: orderService}
}

// CreateOrderHandler gestisce le richieste POST /orders
func (oc *OrderController) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
    // Esempio di creazione di un nuovo ordine dal corpo della richiesta
    var newOrder models.Order
    err := json.NewDecoder(r.Body).Decode(&newOrder)
    if err != nil {
        http.Error(w, "Errore durante la decodifica dei dati dell'ordine", http.StatusBadRequest)
        return
    }

    // Esegui la logica per creare l'ordine utilizzando il servizio OrderService
    err = oc.OrderService.CreateOrder(&newOrder)
    if err != nil {
        http.Error(w, "Errore durante la creazione dell'ordine", http.StatusInternalServerError)
        return
    }

    // Invia una risposta di successo
    w.WriteHeader(http.StatusCreated)
}