package main 

type Products struct {
	id              int   		`json:"id"`
	nome            string  	`json:"nome"`
	riferimento     *string 	`json:"riferimento"`
	categoria       string  	`json:"categoria"`
	prezzotaxescl   float64 	`json:"prezzo_tax_escl"`
	prezzotaxincl   float64 	`json:"prezzo_tax_incl"`
	quantita        int     	`json:"quantita"`
	stato           bool    	`json:"stato"`
	immagine        *string 	`json:"immagine"`
	riepilogo       *string 	`json:"riepilogo"`
	cartaidentita   *string 	`json:"carta_identita"`
	chisono         *string 	`json:"chi_sono"`
	luogodinascita  *string 	`json:"luogo_di_nascita"`
	formazione      *string 	`json:"formazione"`
	carattereestile *string 	`json:"carattere_e_stile"`
	gourmet         *string 	`json:"gourmet"`
	musica          *string 	`json:"musica"`
	cinema          *string 	`json:"cinema"`
	annata          *int    	`json:"annata"`
	premi           *string 	`json:"premi"`
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