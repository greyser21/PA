package models

import "time"

type Advice struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"` // brouillon, publie
	AuthorID  int       `json:"authorId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type Notification struct {
	ID        int       `json:"id"`
	Target    string    `json:"target"` // particulier, professionnel, all
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Channel   string    `json:"channel,omitempty"` // in_app, push
	CreatedAt time.Time `json:"createdAt"`
}

type DashboardStat struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type PDFDocument struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Ref       string    `json:"ref"`
	Path      string    `json:"path"`
	OwnerID   int       `json:"ownerId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type PushSubscription struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	Token    string `json:"token"`
	Platform string `json:"platform"`
	Endpoint string `json:"endpoint,omitempty"`
}

type ForumThread struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ForumPost struct {
	ID        int       `json:"id"`
	ThreadID  int       `json:"threadId"`
	UserID    int       `json:"userId"`
	Content   string    `json:"content"`
	Status    string    `json:"status"` // visible, signale, supprime
	CreatedAt time.Time `json:"createdAt"`
}
