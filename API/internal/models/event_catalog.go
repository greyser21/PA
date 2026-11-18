package models

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Kind        string `json:"kind"` // formation, atelier, evenement
	Date        string `json:"date"`
	Location    string `json:"location"`
	Price       string `json:"price,omitempty"`
	Status      string `json:"status"` // brouillon, en_attente, valide, publie
	Description string `json:"description,omitempty"`
}

type CatalogItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"` // service, formation, evenement
	Price       string `json:"price"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
