package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"` // particulier, professionnel, salarie, admin
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
