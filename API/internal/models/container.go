package models

import "time"

type Container struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Location string `json:"location"`
	Status   string `json:"status"` // disponible, plein, maintenance
}

type Deposit struct {
	ID        int       `json:"id"`
	Item      string    `json:"item"`
	Location  string    `json:"location"`
	Status    string    `json:"status"` // en_attente, valide, refuse, depose, recupere
	OpenCode  string    `json:"openCode,omitempty"`
	Barcode   string    `json:"barcode,omitempty"`
	UserID    int       `json:"userId,omitempty"`
	Container string    `json:"container,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}
