package controllers

import (
    "encoding/json"
    "net/http"

    "yourapp/models"
    "yourapp/services"
)

// ProductController definisce le funzioni per gestire le richieste relative ai prodotti
type ProductController struct {
    ProductService *services.ProductService
}

// NewProductController crea una nuova istanza di ProductController
func NewProductController(productService *services.ProductService) *ProductController {
    return &ProductController{ProductService: productService}
}

// GetProductHandler gestisce le richieste GET /products
func (pc *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {
    products, err := pc.ProductService.GetProducts()
    if err != nil {
        http.Error(w, "Errore durante il recupero dei prodotti", http.StatusInternalServerError)
        return
    }

    // Converti i prodotti in JSON e invia la risposta
    json.NewEncoder(w).Encode(products)
}
