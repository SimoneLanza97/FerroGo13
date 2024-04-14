package main 
import (
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)


type Products struct {
	Id              int   			`json:"id"`
	Nome            string  		`json:"nome"`
	Riferimento     *string 		`json:"riferimento"`
	Categoria       string  		`json:"categoria"`
	Prezzotaxescl   float64 		`json:"prezzo_tax_escl"`
	Prezzotaxincl   float64 		`json:"prezzo_tax_incl"`
	Quantita        int     		`json:"quantita"`
	Stato           bool    		`json:"stato"`
	Immagine        *string 		`json:"immagine"`
	Riepilogo       *string 		`json:"riepilogo"`
	Cartaidentita   *string 		`json:"carta_identita"`
	Chisono         *string 		`json:"chi_sono"`
	Luogodinascita  *string 		`json:"luogo_di_nascita"`
	Formazione      *string 		`json:"formazione"`
	Carattereestile *string 		`json:"carattere_e_stile"`
	Gourmet         *string 		`json:"gourmet"`
	Musica          *string 		`json:"musica"`
	Cinema          *string 		`json:"cinema"`
	Annata          *int    		`json:"annata"`
	Premi           *string 		`json:"premi"`
	CreatedAt		sql.NullTime 	`json:"created_at"`
	UpdatedAt		sql.NullTime 	`json:"updated_at"`
}


type Carts  struct {
    Id          int       		`json:"id"`
    Id_prodotto int       		`json:"id_prodotto"`
    Id_user     int       		`json:"id_user"`
    Quantita    int       		`json:"quantità"`
    Prezzo      float64   		`json:"prezzo"`
	CreatedAt	sql.NullTime 	`json:"created_at"`
	UpdatedAt	sql.NullTime 	`json:"updated_at"`
}


type Users struct {
    Id                      int           	  `json:"id"`
    Nome                    string        	  `json:"nome"`
    Cognome                 string        	  `json:"cognome"`
    Email                   string        	  `json:"email"`
    Password                string        	  `json:"-"`
    Telefono                *int          	  `json:"telefono"`
    Indirizzo_fattura       *string       	  `json:"inidrizzo_fattura"`
    Indirizzo_spedizione    *string       	  `json:"inidrizzo_spedizione"`
    Api_token               *string       	  `json:"api_token"`
    Remember_token          *string       	  `json:"remember_token"`
	CreatedAt				sql.NullTime 	  `json:"created_at"`
	UpdatedAt				sql.NullTime 	  `json:"updated_at"`
}


type CreateUserReq struct {
    Nome                    string        `json:"nome"`
    Cognome                 string        `json:"cognome"`
    Email                   string        `json:"email"`
    Password                string        `json:"password"`

}


type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}


type LoginResponse struct {
	Email  string  `json:"email"`
	Token  string  `json:"token"`
}


func (a *Users) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pw)) == nil
}


func NewUser(nome, cognome, email, password string) (*Users, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
			return nil, err
	}
	return &Users{
		Nome:		nome,
		Cognome:	cognome,
		Email:		email,
		Password:	string(encpw),
	},nil
}