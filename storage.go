package main 

import (
	"database/sql"
	_"github.com/lib/pq"
    "os"
)
type Storage interface {
    // Metodo per recuperare tutti i prodotti
    GetProducts() ([]Products, error)
	// Metodo per recuperare un prodotto dall ID 
	GetProdById(int) (*Products, error)
    // Metodo per creare un nuovo utente
    CreateUser(*Users) error
    // Metodo per autenticare un utente
    AuthenticateUser(username, password string) (bool, error)
}

type PostgresStore struct {
	db *sql.DB
}


func NewPostgresStore() (*PostgresStore, error) {
	
	// variabili valorizzate dalle variabili di ambiente
	user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbname, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// ping per il check al db
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}


func (s *PostgresStore) GetProducts() ([]Products, error) {
    // Query per selezionare tutti i prodotti dalla tabella products
    rows, err := s.db.Query(`SELECT * FROM products`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Slice per memorizzare i prodotti
    products := []Products

    // Itera sui risultati della query
    for rows.Next() {
        product, err := scanintoProd(rows)
        if err != nil {
            return nil, err
        }
        // Aggiungi il prodotto al slice dei prodotti
        products = append(products, product)
    }
    return products, nil
}


func(s *NewPostgresStore) GetProdById(id int) (*Products, error){
	query,err := s.db.Query(`SELECT * FROM products WHERE id = $1`, id)
	if err != nil{
		return nil, err
	}
	for rows.Next() {
		return scanintoProd(rows)
	}
	return nil, fmt.Errorf("product %d not found", id)
}


func(s *NewPostgresStore) CreateUser(user *User) error{
	query := `insert into Users 
	(nome, cognome, email, password)
	values($1, $2, $3, $4)`
	_,err := s.db.Query(
		query,
		user.Nome,
		user.Cognome,
		user.Email,
		user.Password)
	if err != nil {
		return err
	}
	return nil 
}


func scanintoProd(rows *sql.Rows) (*Products, error){
	products := new(Products)
	err := rows.Scan(
		&prod.ID,
        &prod.Nome,
        &prod.Riferimento,
        &prod.Categoria,
        &prod.PrezzoTaxEscl,
        &prod.PrezzoTaxIncl,
        &prod.Quantita,
        &prod.Stato,
        &prod.Immagine,
        &prod.Riepilogo,
        &prod.CartaIdentita,
        &prod.ChiSono,
        &prod.LuogoDiNascita,
        &prod.Formazione,
        &prod.CarattereEstile,
        &prod.Gourmet,
        &prod.Musica,
        &prod.Cinema,
        &prod.Annata,
        &prod.Premi)
        
	return products, nil 
}