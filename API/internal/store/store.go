package store

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"upcycleconnect/api/internal/models"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	mu            sync.RWMutex
	nextID        int
	users         []models.User
	announcements []models.Announcement
	containers    []models.Container
	deposits      []models.Deposit
	projects      []models.Project
	events        []models.Event
	catalog       []models.CatalogItem
	bookings      []models.Booking
	advice        []models.Advice
	notifications []models.Notification
	contracts     []models.Contract
	payments      []models.Payment
	documents     []models.PDFDocument
	pushSubs      []models.PushSubscription
	threads       []models.ForumThread
	posts         []models.ForumPost
}

func New() *Store {
	now := time.Now()
	return &Store{
		nextID: 1000,
		users: []models.User{
			{ID: 1, Name: "Utilisateur Demo", Email: "user@demo.local", Role: "particulier", Active: true, CreatedAt: now},
			{ID: 2, Name: "Atelier Bois", Email: "pro@demo.local", Role: "professionnel", Active: true, CreatedAt: now},
			{ID: 3, Name: "Salarié Demo", Email: "employee@demo.local", Role: "salarie", Active: true, CreatedAt: now},
			{ID: 4, Name: "Admin Demo", Email: "admin@demo.local", Role: "admin", Active: true, CreatedAt: now},
		},
		announcements: []models.Announcement{{ID: 1, Title: "Ancienne commode", Type: "don", Location: "Paris", Status: "validee", OwnerID: 1, CreatedAt: now}, {ID: 2, Title: "Chutes de tissu", Type: "don", Location: "Montreuil", Status: "validee", OwnerID: 1, CreatedAt: now}, {ID: 3, Title: "Lampe à restaurer", Type: "vente", Location: "Paris", Price: "25 €", Status: "en_attente", OwnerID: 1, CreatedAt: now}},
		containers:    []models.Container{{ID: 1, Code: "C-013", Location: "Paris", Status: "disponible"}, {ID: 2, Code: "C-021", Location: "Montreuil", Status: "disponible"}},
		deposits:      []models.Deposit{{ID: 1, Item: "Bois de palettes", Location: "Paris", Status: "depose", Container: "C-013", Barcode: "BC-1001", OpenCode: "OP-4412", UserID: 1, CreatedAt: now}, {ID: 2, Item: "Textile", Location: "Montreuil", Status: "valide", Container: "C-021", Barcode: "BC-1002", OpenCode: "OP-5521", UserID: 1, CreatedAt: now}},
		projects:      []models.Project{{ID: 1, Title: "Collection mobilier", Description: "12 objets", Status: "en_cours", Progress: 68, OwnerID: 2}, {ID: 2, Title: "Lampes recyclées", Description: "8 objets", Status: "termine", Progress: 100, OwnerID: 2}},
		events:        []models.Event{{ID: 1, Title: "Atelier upcycling bois", Kind: "atelier", Date: "Samedi • 14:00", Location: "Paris 10e", Status: "valide"}, {ID: 2, Title: "Formation textile", Kind: "formation", Date: "12 septembre • 18:00", Location: "Montreuil", Status: "valide"}},
		catalog:       []models.CatalogItem{{ID: 1, Title: "Atelier upcycling bois", Category: "atelier", Price: "30 €", Description: "Initiation.", Active: true}, {ID: 2, Title: "Abonnement Pro Premium", Category: "service", Price: "25 €/mois", Description: "Tableaux de bord avancés.", Active: true}},
		bookings:      []models.Booking{{ID: 1, UserID: 1, ItemTitle: "Atelier upcycling bois", Type: "atelier", Date: "Samedi • 14:00", Status: "reserve", CreatedAt: now}},
		advice:        []models.Advice{{ID: 1, Title: "Réparer avant de jeter", Content: "Conseil", Status: "publie", AuthorID: 3, CreatedAt: now}},
		notifications: []models.Notification{{ID: 1, Target: "all", Title: "Bienvenue", Message: "Bienvenue sur UpcycleConnect", Channel: "in_app", CreatedAt: now}},
		contracts:     []models.Contract{{ID: 1, Partner: "Atelier Bois", Type: "abonnement", Amount: "25 €/mois", Status: "actif"}},
	}
}

func (s *Store) nid() int { s.nextID++; return s.nextID }
func (s *Store) Health() map[string]string { return map[string]string{"status": "ok", "service": "upcycleconnect-api"} }

// Users
func (s *Store) ListUsers() []models.User { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.User(nil), s.users...) }
func (s *Store) GetUser(id int) (models.User, error) { s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.users{ if v.ID==id {return v,nil} }; return models.User{}, ErrNotFound }
func (s *Store) CreateUser(v models.User) models.User { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.Role==""{v.Role="particulier"}; v.Active=true; if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; s.users=append(s.users,v); return v }
func (s *Store) UpdateUser(id int,v models.User)(models.User,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.users{ if s.users[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.users[i].CreatedAt}; s.users[i]=v; return v,nil } }; return models.User{},ErrNotFound }
func (s *Store) DeleteUser(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.users{ if v.ID==id{ s.users=append(s.users[:i],s.users[i+1:]...); return nil } }; return ErrNotFound }

// Announcements
func (s *Store) ListAnnouncements() []models.Announcement { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Announcement(nil), s.announcements...) }
func (s *Store) GetAnnouncement(id int)(models.Announcement,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.announcements{ if v.ID==id{return v,nil}}; return models.Announcement{},ErrNotFound }
func (s *Store) CreateAnnouncement(v models.Announcement) models.Announcement { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.Status==""{v.Status="en_attente"}; v.CreatedAt=time.Now(); s.announcements=append([]models.Announcement{v},s.announcements...); return v }
func (s *Store) UpdateAnnouncement(id int,v models.Announcement)(models.Announcement,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.announcements{ if s.announcements[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.announcements[i].CreatedAt}; s.announcements[i]=v; return v,nil } }; return models.Announcement{},ErrNotFound }
func (s *Store) DeleteAnnouncement(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.announcements{ if v.ID==id{ s.announcements=append(s.announcements[:i],s.announcements[i+1:]...); return nil } }; return ErrNotFound }

// Containers
func (s *Store) ListContainers() []models.Container { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Container(nil), s.containers...) }
func (s *Store) GetContainer(id int)(models.Container,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.containers{ if v.ID==id{return v,nil}}; return models.Container{},ErrNotFound }
func (s *Store) CreateContainer(v models.Container) models.Container { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.Status==""{v.Status="disponible"}; s.containers=append(s.containers,v); return v }
func (s *Store) UpdateContainer(id int,v models.Container)(models.Container,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.containers{ if s.containers[i].ID==id{ v.ID=id; s.containers[i]=v; return v,nil } }; return models.Container{},ErrNotFound }
func (s *Store) DeleteContainer(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.containers{ if v.ID==id{ s.containers=append(s.containers[:i],s.containers[i+1:]...); return nil } }; return ErrNotFound }

// Deposits
func (s *Store) ListDeposits() []models.Deposit { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Deposit(nil), s.deposits...) }
func (s *Store) GetDeposit(id int)(models.Deposit,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.deposits{ if v.ID==id{return v,nil}}; return models.Deposit{},ErrNotFound }
func (s *Store) CreateDeposit(v models.Deposit) models.Deposit { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); v.CreatedAt=time.Now(); if v.Status==""{v.Status="en_attente"}; if v.OpenCode==""{v.OpenCode=fmt.Sprintf("OP-%d",1000+v.ID)}; if v.Barcode==""{v.Barcode=fmt.Sprintf("BC-%d",2000+v.ID)}; s.deposits=append([]models.Deposit{v},s.deposits...); return v }
func (s *Store) UpdateDeposit(id int,v models.Deposit)(models.Deposit,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.deposits{ if s.deposits[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.deposits[i].CreatedAt}; s.deposits[i]=v; return v,nil } }; return models.Deposit{},ErrNotFound }
func (s *Store) DeleteDeposit(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.deposits{ if v.ID==id{ s.deposits=append(s.deposits[:i],s.deposits[i+1:]...); return nil } }; return ErrNotFound }

// Projects
func (s *Store) ListProjects() []models.Project { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Project(nil), s.projects...) }
func (s *Store) GetProject(id int)(models.Project,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.projects{ if v.ID==id{return v,nil}}; return models.Project{},ErrNotFound }
func (s *Store) CreateProject(v models.Project) models.Project { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.Status==""{v.Status="en_cours"}; s.projects=append(s.projects,v); return v }
func (s *Store) UpdateProject(id int,v models.Project)(models.Project,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.projects{ if s.projects[i].ID==id{ v.ID=id; s.projects[i]=v; return v,nil } }; return models.Project{},ErrNotFound }
func (s *Store) DeleteProject(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.projects{ if v.ID==id{ s.projects=append(s.projects[:i],s.projects[i+1:]...); return nil } }; return ErrNotFound }

// Events
func (s *Store) ListEvents() []models.Event { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Event(nil), s.events...) }
func (s *Store) GetEvent(id int)(models.Event,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.events{ if v.ID==id{return v,nil}}; return models.Event{},ErrNotFound }
func (s *Store) CreateEvent(v models.Event) models.Event { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.Status==""{v.Status="en_attente"}; s.events=append(s.events,v); return v }
func (s *Store) UpdateEvent(id int,v models.Event)(models.Event,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.events{ if s.events[i].ID==id{ v.ID=id; s.events[i]=v; return v,nil } }; return models.Event{},ErrNotFound }
func (s *Store) DeleteEvent(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.events{ if v.ID==id{ s.events=append(s.events[:i],s.events[i+1:]...); return nil } }; return ErrNotFound }

// Catalog
func (s *Store) ListCatalog() []models.CatalogItem { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.CatalogItem(nil), s.catalog...) }
func (s *Store) GetCatalog(id int)(models.CatalogItem,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.catalog{ if v.ID==id{return v,nil}}; return models.CatalogItem{},ErrNotFound }
func (s *Store) CreateCatalog(v models.CatalogItem) models.CatalogItem { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); v.Active=true; s.catalog=append(s.catalog,v); return v }
func (s *Store) UpdateCatalog(id int,v models.CatalogItem)(models.CatalogItem,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.catalog{ if s.catalog[i].ID==id{ v.ID=id; s.catalog[i]=v; return v,nil } }; return models.CatalogItem{},ErrNotFound }
func (s *Store) DeleteCatalog(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.catalog{ if v.ID==id{ s.catalog=append(s.catalog[:i],s.catalog[i+1:]...); return nil } }; return ErrNotFound }

// Bookings
func (s *Store) ListBookings() []models.Booking { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Booking(nil), s.bookings...) }
func (s *Store) GetBooking(id int)(models.Booking,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.bookings{ if v.ID==id{return v,nil}}; return models.Booking{},ErrNotFound }
func (s *Store) CreateBooking(v models.Booking) models.Booking { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); v.CreatedAt=time.Now(); if v.Status==""{v.Status="reserve"}; s.bookings=append(s.bookings,v); return v }
func (s *Store) UpdateBooking(id int,v models.Booking)(models.Booking,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.bookings{ if s.bookings[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.bookings[i].CreatedAt}; s.bookings[i]=v; return v,nil } }; return models.Booking{},ErrNotFound }
func (s *Store) DeleteBooking(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.bookings{ if v.ID==id{ s.bookings=append(s.bookings[:i],s.bookings[i+1:]...); return nil } }; return ErrNotFound }

// Advice
func (s *Store) ListAdvice() []models.Advice { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Advice(nil), s.advice...) }
func (s *Store) GetAdvice(id int)(models.Advice,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.advice{ if v.ID==id{return v,nil}}; return models.Advice{},ErrNotFound }
func (s *Store) CreateAdvice(v models.Advice) models.Advice { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); v.CreatedAt=time.Now(); if v.Status==""{v.Status="brouillon"}; s.advice=append(s.advice,v); return v }
func (s *Store) UpdateAdvice(id int,v models.Advice)(models.Advice,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.advice{ if s.advice[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.advice[i].CreatedAt}; s.advice[i]=v; return v,nil } }; return models.Advice{},ErrNotFound }
func (s *Store) DeleteAdvice(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.advice{ if v.ID==id{ s.advice=append(s.advice[:i],s.advice[i+1:]...); return nil } }; return ErrNotFound }

// Notifications
func (s *Store) ListNotifications() []models.Notification { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Notification(nil), s.notifications...) }
func (s *Store) GetNotification(id int)(models.Notification,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.notifications{ if v.ID==id{return v,nil}}; return models.Notification{},ErrNotFound }
func (s *Store) CreateNotification(v models.Notification) models.Notification { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); v.CreatedAt=time.Now(); if v.Target==""{v.Target="all"}; s.notifications=append(s.notifications,v); return v }
func (s *Store) DeleteNotification(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.notifications{ if v.ID==id{ s.notifications=append(s.notifications[:i],s.notifications[i+1:]...); return nil } }; return ErrNotFound }

// Contracts
func (s *Store) ListContracts() []models.Contract { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Contract(nil), s.contracts...) }
func (s *Store) GetContract(id int)(models.Contract,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.contracts{ if v.ID==id{return v,nil}}; return models.Contract{},ErrNotFound }
func (s *Store) CreateContract(v models.Contract) models.Contract { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.Status==""{v.Status="en_attente"}; s.contracts=append(s.contracts,v); return v }
func (s *Store) UpdateContract(id int,v models.Contract)(models.Contract,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.contracts{ if s.contracts[i].ID==id{ v.ID=id; s.contracts[i]=v; return v,nil } }; return models.Contract{},ErrNotFound }
func (s *Store) DeleteContract(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.contracts{ if v.ID==id{ s.contracts=append(s.contracts[:i],s.contracts[i+1:]...); return nil } }; return ErrNotFound }

// Dashboards
func (s *Store) DashUser() map[string]any { s.mu.RLock(); defer s.mu.RUnlock(); return map[string]any{"stats":[]models.DashboardStat{{Label:"Annonces publiées",Value:len(s.announcements)},{Label:"Objets déposés",Value:len(s.deposits)},{Label:"Réservations",Value:len(s.bookings)},{Label:"Upcycling Score",Value:84}},"announcements":s.announcements,"activities":s.bookings,"score":84} }
func (s *Store) DashPro() map[string]any { s.mu.RLock(); defer s.mu.RUnlock(); return map[string]any{"stats":[]models.DashboardStat{{Label:"Annonces disponibles",Value:len(s.announcements)},{Label:"Objets/dépôts",Value:len(s.deposits)},{Label:"Projets",Value:len(s.projects)},{Label:"Abonnement",Value:"Premium"}},"items":s.deposits,"projects":s.projects} }
func (s *Store) DashEmployee() map[string]any { s.mu.RLock(); defer s.mu.RUnlock(); return map[string]any{"stats":[]models.DashboardStat{{Label:"Événements",Value:len(s.events)},{Label:"Conseils",Value:len(s.advice)},{Label:"Notifications",Value:len(s.notifications)},{Label:"Contrats",Value:len(s.contracts)}},"activities":s.events,"tasks":s.advice} }
func (s *Store) DashAdmin() map[string]any { s.mu.RLock(); defer s.mu.RUnlock(); pendingA,pendingD:=0,0; for _,a:=range s.announcements{ if a.Status=="en_attente"{pendingA++} }; for _,d:=range s.deposits{ if d.Status=="en_attente"{pendingD++} }; return map[string]any{"stats":[]models.DashboardStat{{Label:"Utilisateurs",Value:len(s.users)},{Label:"Annonces à valider",Value:pendingA},{Label:"Demandes dépôt",Value:pendingD},{Label:"Contrats",Value:len(s.contracts)}},"actions":s.contracts,"activity":s.notifications} }

// Payments
func (s *Store) ListPayments() []models.Payment { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Payment(nil), s.payments...) }
func (s *Store) GetPayment(id int)(models.Payment,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.payments{ if v.ID==id{return v,nil}}; return models.Payment{},ErrNotFound }
func (s *Store) CreatePayment(v models.Payment) models.Payment { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Status==""{v.Status="pending"}; if v.Provider==""{v.Provider="stripe"}; if v.Currency==""{v.Currency="eur"}; s.payments=append(s.payments,v); return v }
func (s *Store) UpdatePayment(id int,v models.Payment)(models.Payment,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.payments{ if s.payments[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.payments[i].CreatedAt}; s.payments[i]=v; return v,nil } }; return models.Payment{},ErrNotFound }
func (s *Store) DeletePayment(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.payments{ if v.ID==id{ s.payments=append(s.payments[:i],s.payments[i+1:]...); return nil } }; return ErrNotFound }

// Documents
func (s *Store) ListDocuments() []models.PDFDocument { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.PDFDocument(nil), s.documents...) }
func (s *Store) GetDocument(id int)(models.PDFDocument,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.documents{ if v.ID==id{return v,nil}}; return models.PDFDocument{},ErrNotFound }
func (s *Store) CreateDocument(v models.PDFDocument) models.PDFDocument { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Path==""{v.Path=fmt.Sprintf("/documents/%s-%d.pdf", v.Type, v.ID)}; s.documents=append(s.documents,v); return v }
func (s *Store) UpdateDocument(id int,v models.PDFDocument)(models.PDFDocument,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.documents{ if s.documents[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.documents[i].CreatedAt}; s.documents[i]=v; return v,nil } }; return models.PDFDocument{},ErrNotFound }
func (s *Store) DeleteDocument(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.documents{ if v.ID==id{ s.documents=append(s.documents[:i],s.documents[i+1:]...); return nil } }; return ErrNotFound }

// Push subscriptions
func (s *Store) ListPushSubs() []models.PushSubscription { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.PushSubscription(nil), s.pushSubs...) }
func (s *Store) GetPushSub(id int)(models.PushSubscription,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.pushSubs{ if v.ID==id{return v,nil}}; return models.PushSubscription{},ErrNotFound }
func (s *Store) CreatePushSub(v models.PushSubscription) models.PushSubscription { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); s.pushSubs=append(s.pushSubs,v); return v }
func (s *Store) DeletePushSub(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.pushSubs{ if v.ID==id{ s.pushSubs=append(s.pushSubs[:i],s.pushSubs[i+1:]...); return nil } }; return ErrNotFound }

// Forum threads
func (s *Store) ListThreads() []models.ForumThread { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.ForumThread(nil), s.threads...) }
func (s *Store) GetThread(id int)(models.ForumThread,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.threads{ if v.ID==id{return v,nil}}; return models.ForumThread{},ErrNotFound }
func (s *Store) CreateThread(v models.ForumThread) models.ForumThread { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Status==""{v.Status="ouvert"}; s.threads=append(s.threads,v); return v }
func (s *Store) UpdateThread(id int,v models.ForumThread)(models.ForumThread,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.threads{ if s.threads[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.threads[i].CreatedAt}; s.threads[i]=v; return v,nil } }; return models.ForumThread{},ErrNotFound }
func (s *Store) DeleteThread(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.threads{ if v.ID==id{ s.threads=append(s.threads[:i],s.threads[i+1:]...); return nil } }; return ErrNotFound }

// Forum posts
func (s *Store) ListPosts() []models.ForumPost { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.ForumPost(nil), s.posts...) }
func (s *Store) GetPost(id int)(models.ForumPost,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.posts{ if v.ID==id{return v,nil}}; return models.ForumPost{},ErrNotFound }
func (s *Store) CreatePost(v models.ForumPost) models.ForumPost { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.nid(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Status==""{v.Status="visible"}; s.posts=append(s.posts,v); return v }
func (s *Store) UpdatePost(id int,v models.ForumPost)(models.ForumPost,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.posts{ if s.posts[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.posts[i].CreatedAt}; s.posts[i]=v; return v,nil } }; return models.ForumPost{},ErrNotFound }
func (s *Store) DeletePost(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.posts{ if v.ID==id{ s.posts=append(s.posts[:i],s.posts[i+1:]...); return nil } }; return ErrNotFound }
