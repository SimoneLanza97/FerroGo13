package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	jwt "github.com/golang-jwt/jwt/v4"
	"log"
	"strconv"
	"os"
)


// struct del server API ha un campo listenAddr che sarà una stringa 
type APIServer struct {
	listenAddr string
	store Storage
}


/* funzione che prende in input una stringa che dirà l'indirizzo al 
quale mettersi in ascolto (es: "127.0.0.1:8080") */ 
func NewAPIServer(listenAddr string,store Storage) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}


// Inizializza il router e associa le routes alle funzioni
func (s *APIServer) Run() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	
	apiRouter.HandleFunc("/products", makeHTTPHandleFunc(s.handleGetProducts)).Methods("GET")
	apiRouter.HandleFunc("/getProdotti/{id}", makeHTTPHandleFunc(s.handleGetProdById)).Methods("GET")
	apiRouter.HandleFunc("/store/user", makeHTTPHandleFunc(s.handleCreateUser)).Methods("POST")
	apiRouter.HandleFunc("/authenticate", makeHTTPHandleFunc(s.handleAuthUser)).Methods("POST")

	// Swagger UI 
	swaggerDir := "/myswagger/" // Assicurati che questo percorso sia corretto
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir(swaggerDir))))

	
	log.Println("JSON API Server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}


// handleGetProducts gets all products
// swagger:operation GET /api/products getProducts
// ---
// summary: Retrieves a list of products.
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/Product'
func (s *APIServer) handleGetProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := s.store.GetProducts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, products)
}

// handleCreateUser creates a new user
// swagger:operation POST /api/store/user createUser
// ---
// summary: Creates a new user.
// consumes:
// - application/json
// parameters:
// - in: body
//   name: user
//   description: The user to create.
//   required: true
//   schema:
//     $ref: '#/definitions/CreateUserReq'
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
//     schema:
//       $ref: '#/definitions/User'
//   '400':
//     description: Invalid user supplied
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error { 
	req := new(CreateUserReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	user, err := NewUser(req.Nome, req.Cognome, req.Email, req.Password)
	if err != nil{
		return err
	}
	if err := s.store.CreateUser(user); err != nil {
		return err 
	}
	return WriteJSON(w, http.StatusOK, user)
}


func createJWT(user *Users) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"accountNumber": user.Email,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}


// handleAuthUser authenticates a user and returns a JWT token
// swagger:operation POST /api/authenticate authenticateUser
// ---
// summary: Authenticates user and returns a JWT token.
// consumes:
// - application/json
// parameters:
// - in: body
//   name: login
//   description: The login details for authentication.
//   required: true
//   schema:
//     $ref: '#/definitions/LoginRequest'
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
//     schema:
//       $ref: '#/definitions/LoginResponse'
//   '401':
//     description: Authentication failed
func (s *APIServer) handleAuthUser(w http.ResponseWriter, r *http.Request) error{
	if r.Method != "POST" {
		return fmt.Errorf("Method is not allowed %s", r.Method)
	}
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err 
	}
	user, err := s.store.GetUserByEmail(string(req.Email))
	if err != nil {
		return err
	}
	if !user.ValidPassword(req.Password){
		return fmt.Errorf("not authenticated")
	}
	token, err := createJWT(user)
	if err != nil {
		return err 
	}
	resp := LoginResponse{
		Email: 	user.Email,
		Token:	token,
	}
	return WriteJSON(w, http.StatusOK, resp)
}


// handleGetProdById gets a single product by id
// swagger:operation GET /api/getProdotti/{id} getProductById
// ---
// summary: Retrieves a product by its ID.
// parameters:
// - name: id
//   in: path
//   description: ID of the product to fetch
//   required: true
//   type: integer
//   format: int64
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
//     schema:
//       $ref: '#/definitions/Product'
//   '400':
//     description: Invalid ID supplied
//   '404':
//     description: Product not found
func (s *APIServer) handleGetProdById(w http.ResponseWriter, r *http.Request) error{
	if r.Method == "GET" {
		id,err := getID(r)
		if err != nil {
			return err
		}
		product, err := s.store.GetProdById(id)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, product)
	}
	return fmt.Errorf("method is not allowed %s", r.Method)
}


//funzione che prende in input la funzione responsewriter , status e 
// v (il corpo della risposta) 
func WriteJSON(w http.ResponseWriter, status int , v any) error{
	// setta l'headet Content-Type application/json 
	w.Header().Add("Content-Type", "application/json")
	// writeHeader riporta l'Header con lo status
	w.WriteHeader(status)
	// ritorna la codifica json di v 
	return json.NewEncoder(w).Encode(v)
}


// Firma le Apifunc come funzioni che prendo in innput http.ReponseWriter come 
// puntatore di http.Request 
type apiFunc func(http.ResponseWriter, *http.Request) error 
// tipo ApiError ha un campo Error in formato json 
type ApiError struct {
	Error string
}


/* prende in input apiFunc e restituisce http.HandlerFunc http.HandlerFunc verifica
 la richiesta e la risposta e gestisce l'errore nel caso ci sia */ 
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}
	}
}


func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}
