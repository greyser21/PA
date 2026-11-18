package models

import "time"

// Generic API response helpers can be built in handlers; models define domain entities.

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"` // particulier, professionnel, salarie, admin
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

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

type Container struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Location string `json:"location"`
	Status   string `json:"status"` // disponible, plein, maintenance
}

type Deposit struct {
	ID         int       `json:"id"`
	Item       string    `json:"item"`
	Location   string    `json:"location"`
	Status     string    `json:"status"` // en_attente, valide, refuse, depose, recupere
	OpenCode   string    `json:"openCode,omitempty"`
	Barcode    string    `json:"barcode,omitempty"`
	UserID     int       `json:"userId,omitempty"`
	Container  string    `json:"container,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // en_cours, termine
	Progress    int    `json:"progress"`
	OwnerID     int    `json:"ownerId,omitempty"`
}

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

type Booking struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId,omitempty"`
	ItemID    int       `json:"itemId,omitempty"`
	ItemTitle string    `json:"itemTitle"`
	Type      string    `json:"type"` // formation, service, evenement
	Date      string    `json:"date"`
	Status    string    `json:"status"` // reserve, paye, annule
	CreatedAt time.Time `json:"createdAt"`
}

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
	CreatedAt time.Time `json:"createdAt"`
}

type Contract struct {
	ID       int    `json:"id"`
	Partner  string `json:"partner"`
	Type     string `json:"type"` // abonnement, publicite
	Amount   string `json:"amount"`
	Status   string `json:"status"` // actif, en_attente, resilie
	Details  string `json:"details,omitempty"`
}

// Dashboard view models
type Stat struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type ActivityView struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

type TaskView struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  string `json:"status"`
}

type RecentView struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	When    string `json:"when"`
}

type ContainerItemView struct {
	ID        int    `json:"id"`
	Material  string `json:"material"`
	Container string `json:"container"`
	Location  string `json:"location"`
	Status    string `json:"status"`
}

type UserDashboard struct {
	Stats         []Stat         `json:"stats"`
	Announcements []Announcement `json:"announcements"`
	Activities    []ActivityView `json:"activities"`
	Score         int            `json:"score"`
}

type ProDashboard struct {
	Stats    []Stat              `json:"stats"`
	Items    []ContainerItemView `json:"items"`
	Projects []Project           `json:"projects"`
}

type EmployeeDashboard struct {
	Stats      []Stat         `json:"stats"`
	Activities []ActivityView `json:"activities"`
	Tasks      []TaskView     `json:"tasks"`
}

type AdminDashboard struct {
	Stats    []Stat       `json:"stats"`
	Actions  []TaskView   `json:"actions"`
	Activity []RecentView `json:"activity"`
}

// Additional entities for extended backend coverage
type Payment struct {
	ID           int       `json:"id"`
	BookingID    int       `json:"bookingId,omitempty"`
	ContractID   int       `json:"contractId,omitempty"`
	Amount       string    `json:"amount"`
	Currency     string    `json:"currency"`
	Status       string    `json:"status"`
	Provider     string    `json:"provider"`
	StripeIntent string    `json:"stripeIntent,omitempty"`
	PDFPath      string    `json:"pdfPath,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
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
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Token     string `json:"token"`
	Platform  string `json:"platform"`
	Endpoint  string `json:"endpoint"`
}

type ForumPost struct {
	ID        int       `json:"id"`
	ThreadID  int       `json:"threadId"`
	UserID    int       `json:"userId"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ForumThread struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
