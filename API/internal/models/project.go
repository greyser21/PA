package models

type Project struct {
	ID                   int    `json:"id"`
	Title                string `json:"title"`
	Description          string `json:"description"`
	Status               string `json:"status"` // en_cours, termine
	Progress             int    `json:"progress"`
	OwnerID              int    `json:"ownerId,omitempty"`
	UpcyclingScoreImpact int    `json:"upcyclingScoreImpact,omitempty"`
}
