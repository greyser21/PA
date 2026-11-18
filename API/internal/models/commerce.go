package models

import "time"

type Booking struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId,omitempty"`
	ItemID    int       `json:"itemId,omitempty"`
	ItemTitle string    `json:"itemTitle"`
	Type      string    `json:"type"`
	Date      string    `json:"date"`
	Status    string    `json:"status"` // reserve, paye, annule
	CreatedAt time.Time `json:"createdAt"`
}

type Contract struct {
	ID      int    `json:"id"`
	Partner string `json:"partner"`
	Type    string `json:"type"` // abonnement, publicite
	Amount  string `json:"amount"`
	Status  string `json:"status"` // actif, en_attente, resilie
	Details string `json:"details,omitempty"`
}

type Payment struct {
	ID        int       `json:"id"`
	BookingID int       `json:"bookingId,omitempty"`
	Amount    string    `json:"amount"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"` // pending, succeeded, failed
	Provider  string    `json:"provider"` // stripe
	CreatedAt time.Time `json:"createdAt"`
}
