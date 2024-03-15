// models/order.go

package models

// Order rappresenta un ordine effettuato da un utente
type Order struct {
    ID         int
    UserID     int    // ID dell'utente che ha effettuato l'ordine
    ProductID  int    // ID del prodotto ordinato
    Quantity   int    // Quantità del prodotto ordinato
    TotalPrice float64
    // Altri campi ordine come data, stato, ecc.
}
