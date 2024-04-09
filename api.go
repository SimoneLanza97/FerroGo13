package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)


// struct del server API ha un campo listenAddr che sarà una stringa 
type APIServer struct {
	listenAddr string
}


/* funzione che prende in input una stringa che dirà l'indirizzo al 
quale mettersi in ascolto (es: "127.0.0.1:8080") */ 
func NewAPIServer(listenAddr string) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
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
	
	log.Println("JSON API Server running on port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}


func (s *APIServer) handleGetProducts(w http.ResponseWriter, r *http.Request) error{
	return nil
}
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error{
	return nil
}
func (s *APIServer) handleAuthUser(w http.ResponseWriter, r *http.Request) error{
	return nil
}
func (s *APIServer) handleGetProdById(w http.ResponseWriter, r *http.Request) error{
	return nil
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


// func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error{
// 	if r.Method == "GET" {
// 		return s.handleGetAccount(w, r) 	
// 	}
// 	if r.Method == "POST" {
// 		return s.handleCreateAccount(w, r) 	
// 	}
// 	if r.Method == "DELETE" {
// 		return s.handleDeleteAccount(w, r) 	
// 	}
// 	return  fmt.Errorf("method not allowed %s", r.Method)
// }