package main 

type Products struct {
	Id              int   		`json:"id"`
	Nome            string  	`json:"nome"`
	Riferimento     *string 	`json:"riferimento"`
	Categoria       string  	`json:"categoria"`
	Prezzotaxescl   float64 	`json:"prezzo_tax_escl"`
	Prezzotaxincl   float64 	`json:"prezzo_tax_incl"`
	Quantita        int     	`json:"quantita"`
	Stato           bool    	`json:"stato"`
	Immagine        *string 	`json:"immagine"`
	Riepilogo       *string 	`json:"riepilogo"`
	Cartaidentita   *string 	`json:"carta_identita"`
	Chisono         *string 	`json:"chi_sono"`
	Luogodinascita  *string 	`json:"luogo_di_nascita"`
	Formazione      *string 	`json:"formazione"`
	Carattereestile *string 	`json:"carattere_e_stile"`
	Gourmet         *string 	`json:"gourmet"`
	Musica          *string 	`json:"musica"`
	Cinema          *string 	`json:"cinema"`
	Annata          *int    	`json:"annata"`
	Premi           *string 	`json:"premi"`
}
type Carts  struct {
    ID          int      `json:"id"`
    Id_prodotto int      `json:"id_prodotto"`
    Id_user     int      `json:"id_user"`
    Quantita    int      `json:"quantità"`
    Prezzo      float64  `json:"prezzo"`
}
type Users struct {
    ID                      int           `json:"id"`
    Nome                    string        `json:"nome"`
    Cognome                 string        `json:"cognome"`
    Email                   string        `json:"email"`
    Password                string        `json:"password"`
    Telefono                *int          `json:"telefono"`
    Indirizzo_fattura       *string       `json:"inidrizzo_fattura"`
    Indirizzo_spedizione    *string       `json:"inidrizzo_spedizione"`
    Api_token               *string       `json:"api_token"`
    Remember_token          *string       `json:"remember_token"`
}

type CreateUserReq struct {
    Nome                    string        `json:"nome"`
    Cognome                 string        `json:"cognome"`
    Email                   string        `json:"email"`
    Password                string        `json:"password"`

}

func NewUser(nome, cognome, email, password) (*Users, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
			return nil, err
	}
	return &Users{
		Nome:		Nome,
		Cognome:	Cognome,
		Email:		Email,
		Password:	string(encpw)
	},nil
}