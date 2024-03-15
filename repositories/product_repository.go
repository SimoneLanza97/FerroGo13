package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    // Connessione al database PostgreSQL
    db, err := sql.Open("postgres", "postgres://user:password@localhost/database_name?sslmode=disable")
    if err != nil {
        log.Fatal("Errore durante la connessione al database:", err)
    }
    defer db.Close()