package models

import "time"

type Announcement struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Type      string    `json:"type"` // don, vente
	Location  string    `json:"location"`
	Price     string    `json:"price,omitempty"`
	Status    string    `json:"status"` // en_attente, validee, refusee
	OwnerID   int       `json:"ownerId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}
