package store

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"upcycleconnect/internal/models"
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
		announcements: []models.Announcement{
			{ID: 1, Title: "Ancienne commode", Type: "don", Location: "Paris", Status: "validee", OwnerID: 1, CreatedAt: now},
			{ID: 2, Title: "Chutes de tissu", Type: "don", Location: "Montreuil", Status: "validee", OwnerID: 1, CreatedAt: now},
			{ID: 3, Title: "Lampe à restaurer", Type: "vente", Location: "Paris", Price: "25 €", Status: "en_attente", OwnerID: 1, CreatedAt: now},
		},
		containers: []models.Container{
			{ID: 1, Code: "C-013", Location: "Paris", Status: "disponible"},
			{ID: 2, Code: "C-021", Location: "Montreuil", Status: "disponible"},
			{ID: 3, Code: "C-008", Location: "Ivry", Status: "disponible"},
		},
		deposits: []models.Deposit{
			{ID: 1, Item: "Bois de palettes", Location: "Paris", Status: "depose", Container: "C-013", Barcode: "BC-1001", OpenCode: "OP-4412", UserID: 1, CreatedAt: now},
			{ID: 2, Item: "Textile", Location: "Montreuil", Status: "depose", Container: "C-021", Barcode: "BC-1002", OpenCode: "OP-5521", UserID: 1, CreatedAt: now},
			{ID: 3, Item: "Éléments métalliques", Location: "Ivry", Status: "valide", Container: "C-008", Barcode: "BC-1003", OpenCode: "OP-7781", UserID: 1, CreatedAt: now},
		},
		projects: []models.Project{
			{ID: 1, Title: "Collection mobilier", Description: "12 objets", Status: "en_cours", Progress: 68, OwnerID: 2},
			{ID: 2, Title: "Lampes recyclées", Description: "8 objets", Status: "termine", Progress: 100, OwnerID: 2},
		},
		events: []models.Event{
			{ID: 1, Title: "Atelier upcycling bois", Kind: "atelier", Date: "Samedi • 14:00", Location: "Paris 10e", Status: "valide"},
			{ID: 2, Title: "Formation textile", Kind: "formation", Date: "12 septembre • 18:00", Location: "Montreuil", Status: "valide"},
			{ID: 3, Title: "Formation « Bois »", Kind: "formation", Date: "À planifier", Location: "Paris", Status: "en_attente"},
		},
		catalog: []models.CatalogItem{
			{ID: 1, Title: "Atelier upcycling bois", Category: "atelier", Price: "30 €", Description: "Initiation à la transformation du bois.", Active: true},
			{ID: 2, Title: "Formation textile", Category: "formation", Price: "50 €", Description: "Réparer et transformer le textile.", Active: true},
			{ID: 3, Title: "Abonnement Pro Premium", Category: "service", Price: "25 €/mois", Description: "Tableaux de bord, alertes, stats avancées.", Active: true},
		},
		bookings: []models.Booking{
			{ID: 1, UserID: 1, ItemID: 1, ItemTitle: "Atelier upcycling bois", Type: "atelier", Date: "Samedi • 14:00", Status: "reserve", CreatedAt: now},
		},
		advice: []models.Advice{
			{ID: 1, Title: "Réparer avant de jeter", Content: "Quelques gestes simples pour prolonger la vie des objets.", Status: "publie", AuthorID: 3, CreatedAt: now},
			{ID: 2, Title: "Conseil n°18", Content: "Brouillon de conseil à valider.", Status: "brouillon", AuthorID: 3, CreatedAt: now},
		},
		notifications: []models.Notification{
			{ID: 1, Target: "particulier", Title: "Bienvenue", Message: "Bienvenue sur UpcycleConnect.", CreatedAt: now},
		},
		contracts: []models.Contract{
			{ID: 1, Partner: "Atelier Bois", Type: "abonnement", Amount: "25 €/mois", Status: "actif", Details: "Premium"},
		},
	}
}

func (s *Store) id() int { s.nextID++; return s.nextID }

func (s *Store) Health() map[string]string {
	return map[string]string{"status": "ok", "service": "upcycleconnect-api"}
}

// --- Dashboards ---
func (s *Store) UserDashboard() models.UserDashboard {
	s.mu.RLock(); defer s.mu.RUnlock()
	ann := append([]models.Announcement(nil), s.announcements...)
	acts := []models.ActivityView{}
	for _, b := range s.bookings {
		acts = append(acts, models.ActivityView{ID: b.ID, Title: b.ItemTitle, Date: b.Date, Status: b.Status})
	}
	for _, e := range s.events {
		if e.Status == "valide" {
			acts = append(acts, models.ActivityView{ID: 100 + e.ID, Title: e.Title, Date: e.Date, Status: "À venir"})
		}
	}
	return models.UserDashboard{
		Stats: []models.Stat{
			{Label: "Annonces publiées", Value: len(ann)},
			{Label: "Objets déposés", Value: countDeposits(s.deposits)},
			{Label: "Réservations", Value: len(s.bookings)},
			{Label: "Upcycling Score", Value: 84},
		},
		Announcements: ann,
		Activities:    acts,
		Score:         84,
	}
}
func countDeposits(ds []models.Deposit) int { c := 0; for _, d := range ds { if d.Status != "refuse" { c++ } }; return c }

func (s *Store) ProDashboard() models.ProDashboard {
	s.mu.RLock(); defer s.mu.RUnlock()
	items := []models.ContainerItemView{}
	for _, d := range s.deposits {
		if d.Status == "depose" || d.Status == "valide" {
			st := "Disponible"; if d.Status == "valide" { st = "Réservé" }
			items = append(items, models.ContainerItemView{ID: d.ID, Material: d.Item, Container: d.Container, Location: d.Location, Status: st})
		}
	}
	return models.ProDashboard{
		Stats: []models.Stat{
			{Label: "Annonces disponibles", Value: len(s.announcements)},
			{Label: "Objets réservés", Value: len(items)},
			{Label: "Projets", Value: len(s.projects)},
			{Label: "Abonnement", Value: "Premium"},
		},
		Items: items, Projects: append([]models.Project(nil), s.projects...),
	}
}

func (s *Store) EmployeeDashboard() models.EmployeeDashboard {
	s.mu.RLock(); defer s.mu.RUnlock()
	acts := []models.ActivityView{}
	for _, e := range s.events { acts = append(acts, models.ActivityView{ID: e.ID, Title: e.Title, Date: e.Date, Status: e.Status}) }
	tasks := []models.TaskView{}
	for _, e := range s.events { if e.Status == "en_attente" { tasks = append(tasks, models.TaskView{ID: e.ID, Title: e.Title, Details: "Événement/formation à valider", Status: "Validation"}) } }
	for _, a := range s.advice { if a.Status == "brouillon" { tasks = append(tasks, models.TaskView{ID: 100 + a.ID, Title: a.Title, Details: "Brouillon à publier", Status: "Brouillon"}) } }
	tasks = append(tasks, models.TaskView{ID: 999, Title: "Forum — modération", Details: "Signalements à traiter", Status: "À traiter"})
	return models.EmployeeDashboard{
		Stats: []models.Stat{{Label: "Formations", Value: countKind(s.events, "formation")}, {Label: "Événements", Value: len(s.events)}, {Label: "Conseils publiés", Value: countPublished(s.advice)}, {Label: "À traiter", Value: len(tasks)}},
		Activities: acts, Tasks: tasks,
	}
}
func countKind(es []models.Event, k string) int { c:=0; for _,e:=range es{ if e.Kind==k{c++} }; return c }
func countPublished(as []models.Advice) int { c:=0; for _,a:=range as{ if a.Status=="publie"{c++} }; return c }

func (s *Store) AdminDashboard() models.AdminDashboard {
	s.mu.RLock(); defer s.mu.RUnlock()
	pendingAnn := 0; for _,a:=range s.announcements{ if a.Status=="en_attente"{pendingAnn++} }
	pendingDep := 0; for _,d:=range s.deposits{ if d.Status=="en_attente"{pendingDep++} }
	actions := []models.TaskView{
		{ID: 1, Title: fmt.Sprintf("%d annonces", pendingAnn), Details: "Annonces en attente de validation", Status: "À valider"},
		{ID: 2, Title: fmt.Sprintf("%d dépôts", pendingDep), Details: "Demandes de conteneurs", Status: "À traiter"},
		{ID: 3, Title: fmt.Sprintf("%d événements", len(s.events)), Details: "Événements au catalogue", Status: "Suivi"},
		{ID: 4, Title: fmt.Sprintf("%d contrats", len(s.contracts)), Details: "Contrats & abonnements", Status: "Gestion"},
	}
	activity := []models.RecentView{}
	for i := len(s.notifications)-1; i >= 0 && len(activity) < 5; i-- {
		n := s.notifications[i]
		activity = append(activity, models.RecentView{ID: n.ID, Title: n.Title, Details: n.Message, When: n.CreatedAt.Format("02/01/2006 15:04")})
	}
	if len(activity)==0 { activity = append(activity, models.RecentView{ID:1, Title:"Activité", Details:"Aucune notification récente", When:"-"}) }
	return models.AdminDashboard{
		Stats: []models.Stat{{Label: "Utilisateurs", Value: len(s.users)}, {Label: "Annonces à valider", Value: pendingAnn}, {Label: "Demandes dépôt", Value: pendingDep}, {Label: "Contrats", Value: len(s.contracts)}},
		Actions: actions, Activity: activity,
	}
}

// --- Generic helpers ---
func matchID(id int, target int) bool { return id == target }

// Users CRUD
func (s *Store) ListUsers() []models.User { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.User(nil), s.users...) }
func (s *Store) GetUser(id int) (models.User, error) { s.mu.RLock(); defer s.mu.RUnlock(); for _,u:=range s.users{ if u.ID==id {return u,nil} }; return models.User{}, ErrNotFound }
func (s *Store) CreateUser(u models.User) models.User { s.mu.Lock(); defer s.mu.Unlock(); u.ID=s.id(); if u.CreatedAt.IsZero(){u.CreatedAt=time.Now()}; if u.Role==""{u.Role="particulier"}; u.Active=true; s.users=append(s.users,u); return u }
func (s *Store) UpdateUser(id int, u models.User) (models.User,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.users{ if s.users[i].ID==id { u.ID=id; if u.CreatedAt.IsZero(){u.CreatedAt=s.users[i].CreatedAt}; s.users[i]=u; return u,nil } }; return models.User{},ErrNotFound }
func (s *Store) DeleteUser(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,u:=range s.users{ if u.ID==id { s.users=append(s.users[:i],s.users[i+1:]...); return nil } }; return ErrNotFound }

// Announcements CRUD
func (s *Store) ListAnnouncements() []models.Announcement { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Announcement(nil), s.announcements...) }
func (s *Store) GetAnnouncement(id int)(models.Announcement,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,a:=range s.announcements{ if a.ID==id{return a,nil} }; return models.Announcement{},ErrNotFound }
func (s *Store) CreateAnnouncement(a models.Announcement) models.Announcement { s.mu.Lock(); defer s.mu.Unlock(); a.ID=s.id(); if a.Status==""{a.Status="en_attente"}; a.CreatedAt=time.Now(); s.announcements=append([]models.Announcement{a},s.announcements...); s.notifications=append(s.notifications, models.Notification{ID:s.id(),Target:"admin",Title:"Nouvelle annonce",Message:a.Title,CreatedAt:time.Now()}); return a }
func (s *Store) UpdateAnnouncement(id int,a models.Announcement)(models.Announcement,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.announcements{ if s.announcements[i].ID==id{ a.ID=id; if a.CreatedAt.IsZero(){a.CreatedAt=s.announcements[i].CreatedAt}; s.announcements[i]=a; return a,nil } }; return models.Announcement{},ErrNotFound }
func (s *Store) DeleteAnnouncement(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,a:=range s.announcements{ if a.ID==id{ s.announcements=append(s.announcements[:i],s.announcements[i+1:]...); return nil } }; return ErrNotFound }

// Containers CRUD
func (s *Store) ListContainers() []models.Container { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Container(nil), s.containers...) }
func (s *Store) GetContainer(id int)(models.Container,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,c:=range s.containers{ if c.ID==id{return c,nil}}; return models.Container{},ErrNotFound }
func (s *Store) CreateContainer(c models.Container) models.Container { s.mu.Lock(); defer s.mu.Unlock(); c.ID=s.id(); if c.Status==""{c.Status="disponible"}; s.containers=append(s.containers,c); return c }
func (s *Store) UpdateContainer(id int,c models.Container)(models.Container,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.containers{ if s.containers[i].ID==id{ c.ID=id; s.containers[i]=c; return c,nil } }; return models.Container{},ErrNotFound }
func (s *Store) DeleteContainer(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,c:=range s.containers{ if c.ID==id{ s.containers=append(s.containers[:i],s.containers[i+1:]...); return nil } }; return ErrNotFound }

// Deposits CRUD
func (s *Store) ListDeposits() []models.Deposit { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Deposit(nil), s.deposits...) }
func (s *Store) GetDeposit(id int)(models.Deposit,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,d:=range s.deposits{ if d.ID==id{return d,nil}}; return models.Deposit{},ErrNotFound }
func (s *Store) CreateDeposit(d models.Deposit) models.Deposit { s.mu.Lock(); defer s.mu.Unlock(); d.ID=s.id(); d.CreatedAt=time.Now(); if d.Status==""{d.Status="en_attente"}; if d.OpenCode==""{d.OpenCode=fmt.Sprintf("OP-%d", 1000+d.ID)}; if d.Barcode==""{d.Barcode=fmt.Sprintf("BC-%d", 2000+d.ID)}; s.deposits=append([]models.Deposit{d},s.deposits...); s.notifications=append(s.notifications, models.Notification{ID:s.id(),Target:"admin",Title:"Demande de dépôt",Message:d.Item,CreatedAt:time.Now()}); return d }
func (s *Store) UpdateDeposit(id int,d models.Deposit)(models.Deposit,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.deposits{ if s.deposits[i].ID==id{ d.ID=id; if d.CreatedAt.IsZero(){d.CreatedAt=s.deposits[i].CreatedAt}; s.deposits[i]=d; return d,nil } }; return models.Deposit{},ErrNotFound }
func (s *Store) DeleteDeposit(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,d:=range s.deposits{ if d.ID==id{ s.deposits=append(s.deposits[:i],s.deposits[i+1:]...); return nil } }; return ErrNotFound }

// Projects CRUD
func (s *Store) ListProjects() []models.Project { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Project(nil), s.projects...) }
func (s *Store) GetProject(id int)(models.Project,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,p:=range s.projects{ if p.ID==id{return p,nil}}; return models.Project{},ErrNotFound }
func (s *Store) CreateProject(p models.Project) models.Project { s.mu.Lock(); defer s.mu.Unlock(); p.ID=s.id(); if p.Status==""{p.Status="en_cours"}; s.projects=append(s.projects,p); return p }
func (s *Store) UpdateProject(id int,p models.Project)(models.Project,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.projects{ if s.projects[i].ID==id{ p.ID=id; s.projects[i]=p; return p,nil } }; return models.Project{},ErrNotFound }
func (s *Store) DeleteProject(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,p:=range s.projects{ if p.ID==id{ s.projects=append(s.projects[:i],s.projects[i+1:]...); return nil } }; return ErrNotFound }

// Events CRUD
func (s *Store) ListEvents() []models.Event { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Event(nil), s.events...) }
func (s *Store) GetEvent(id int)(models.Event,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,e:=range s.events{ if e.ID==id{return e,nil}}; return models.Event{},ErrNotFound }
func (s *Store) CreateEvent(e models.Event) models.Event { s.mu.Lock(); defer s.mu.Unlock(); e.ID=s.id(); if e.Status==""{e.Status="en_attente"}; s.events=append(s.events,e); return e }
func (s *Store) UpdateEvent(id int,e models.Event)(models.Event,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.events{ if s.events[i].ID==id{ e.ID=id; s.events[i]=e; return e,nil } }; return models.Event{},ErrNotFound }
func (s *Store) DeleteEvent(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,e:=range s.events{ if e.ID==id{ s.events=append(s.events[:i],s.events[i+1:]...); return nil } }; return ErrNotFound }

// Catalog CRUD
func (s *Store) ListCatalog() []models.CatalogItem { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.CatalogItem(nil), s.catalog...) }
func (s *Store) GetCatalogItem(id int)(models.CatalogItem,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,c:=range s.catalog{ if c.ID==id{return c,nil}}; return models.CatalogItem{},ErrNotFound }
func (s *Store) CreateCatalogItem(c models.CatalogItem) models.CatalogItem { s.mu.Lock(); defer s.mu.Unlock(); c.ID=s.id(); c.Active=true; s.catalog=append(s.catalog,c); return c }
func (s *Store) UpdateCatalogItem(id int,c models.CatalogItem)(models.CatalogItem,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.catalog{ if s.catalog[i].ID==id{ c.ID=id; s.catalog[i]=c; return c,nil } }; return models.CatalogItem{},ErrNotFound }
func (s *Store) DeleteCatalogItem(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,c:=range s.catalog{ if c.ID==id{ s.catalog=append(s.catalog[:i],s.catalog[i+1:]...); return nil } }; return ErrNotFound }

// Bookings CRUD
func (s *Store) ListBookings() []models.Booking { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Booking(nil), s.bookings...) }
func (s *Store) GetBooking(id int)(models.Booking,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,b:=range s.bookings{ if b.ID==id{return b,nil}}; return models.Booking{},ErrNotFound }
func (s *Store) CreateBooking(b models.Booking) models.Booking { s.mu.Lock(); defer s.mu.Unlock(); b.ID=s.id(); b.CreatedAt=time.Now(); if b.Status==""{b.Status="reserve"}; s.bookings=append(s.bookings,b); return b }
func (s *Store) UpdateBooking(id int,b models.Booking)(models.Booking,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.bookings{ if s.bookings[i].ID==id{ b.ID=id; if b.CreatedAt.IsZero(){b.CreatedAt=s.bookings[i].CreatedAt}; s.bookings[i]=b; return b,nil } }; return models.Booking{},ErrNotFound }
func (s *Store) DeleteBooking(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,b:=range s.bookings{ if b.ID==id{ s.bookings=append(s.bookings[:i],s.bookings[i+1:]...); return nil } }; return ErrNotFound }

// Advice CRUD
func (s *Store) ListAdvice() []models.Advice { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Advice(nil), s.advice...) }
func (s *Store) GetAdvice(id int)(models.Advice,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,a:=range s.advice{ if a.ID==id{return a,nil}}; return models.Advice{},ErrNotFound }
func (s *Store) CreateAdvice(a models.Advice) models.Advice { s.mu.Lock(); defer s.mu.Unlock(); a.ID=s.id(); a.CreatedAt=time.Now(); if a.Status==""{a.Status="brouillon"}; s.advice=append(s.advice,a); return a }
func (s *Store) UpdateAdvice(id int,a models.Advice)(models.Advice,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.advice{ if s.advice[i].ID==id{ a.ID=id; if a.CreatedAt.IsZero(){a.CreatedAt=s.advice[i].CreatedAt}; s.advice[i]=a; return a,nil } }; return models.Advice{},ErrNotFound }
func (s *Store) DeleteAdvice(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,a:=range s.advice{ if a.ID==id{ s.advice=append(s.advice[:i],s.advice[i+1:]...); return nil } }; return ErrNotFound }

// Notifications CRUD
func (s *Store) ListNotifications() []models.Notification { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Notification(nil), s.notifications...) }
func (s *Store) GetNotification(id int)(models.Notification,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,n:=range s.notifications{ if n.ID==id{return n,nil}}; return models.Notification{},ErrNotFound }
func (s *Store) CreateNotification(n models.Notification) models.Notification { s.mu.Lock(); defer s.mu.Unlock(); n.ID=s.id(); n.CreatedAt=time.Now(); if strings.TrimSpace(n.Target)==""{n.Target="all"}; s.notifications=append(s.notifications,n); return n }
func (s *Store) DeleteNotification(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,n:=range s.notifications{ if n.ID==id{ s.notifications=append(s.notifications[:i],s.notifications[i+1:]...); return nil } }; return ErrNotFound }

// Contracts CRUD
func (s *Store) ListContracts() []models.Contract { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Contract(nil), s.contracts...) }
func (s *Store) GetContract(id int)(models.Contract,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,c:=range s.contracts{ if c.ID==id{return c,nil}}; return models.Contract{},ErrNotFound }
func (s *Store) CreateContract(c models.Contract) models.Contract { s.mu.Lock(); defer s.mu.Unlock(); c.ID=s.id(); if c.Status==""{c.Status="en_attente"}; s.contracts=append(s.contracts,c); return c }
func (s *Store) UpdateContract(id int,c models.Contract)(models.Contract,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.contracts{ if s.contracts[i].ID==id{ c.ID=id; s.contracts[i]=c; return c,nil } }; return models.Contract{},ErrNotFound }
func (s *Store) DeleteContract(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,c:=range s.contracts{ if c.ID==id{ s.contracts=append(s.contracts[:i],s.contracts[i+1:]...); return nil } }; return ErrNotFound }

// Payments CRUD
func (s *Store) ListPayments() []models.Payment { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.Payment(nil), s.payments...) }
func (s *Store) GetPayment(id int)(models.Payment,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.payments{ if v.ID==id{return v,nil}}; return models.Payment{},ErrNotFound }
func (s *Store) CreatePayment(v models.Payment) models.Payment { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.id(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Status==""{v.Status="pending"}; if v.Provider==""{v.Provider="stripe"}; if v.Currency==""{v.Currency="eur"}; s.payments=append(s.payments,v); return v }
func (s *Store) UpdatePayment(id int,v models.Payment)(models.Payment,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.payments{ if s.payments[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.payments[i].CreatedAt}; s.payments[i]=v; return v,nil } }; return models.Payment{},ErrNotFound }
func (s *Store) DeletePayment(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.payments{ if v.ID==id{ s.payments=append(s.payments[:i],s.payments[i+1:]...); return nil } }; return ErrNotFound }

// PDF Documents CRUD
func (s *Store) ListDocuments() []models.PDFDocument { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.PDFDocument(nil), s.documents...) }
func (s *Store) GetDocument(id int)(models.PDFDocument,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.documents{ if v.ID==id{return v,nil}}; return models.PDFDocument{},ErrNotFound }
func (s *Store) CreateDocument(v models.PDFDocument) models.PDFDocument { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.id(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Path==""{v.Path=fmt.Sprintf("/documents/%s-%d.pdf", v.Type, v.ID)}; s.documents=append(s.documents,v); return v }
func (s *Store) UpdateDocument(id int,v models.PDFDocument)(models.PDFDocument,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.documents{ if s.documents[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.documents[i].CreatedAt}; s.documents[i]=v; return v,nil } }; return models.PDFDocument{},ErrNotFound }
func (s *Store) DeleteDocument(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.documents{ if v.ID==id{ s.documents=append(s.documents[:i],s.documents[i+1:]...); return nil } }; return ErrNotFound }

// Push subscriptions CRUD
func (s *Store) ListPushSubs() []models.PushSubscription { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.PushSubscription(nil), s.pushSubs...) }
func (s *Store) GetPushSub(id int)(models.PushSubscription,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.pushSubs{ if v.ID==id{return v,nil}}; return models.PushSubscription{},ErrNotFound }
func (s *Store) CreatePushSub(v models.PushSubscription) models.PushSubscription { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.id(); s.pushSubs=append(s.pushSubs,v); return v }
func (s *Store) DeletePushSub(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.pushSubs{ if v.ID==id{ s.pushSubs=append(s.pushSubs[:i],s.pushSubs[i+1:]...); return nil } }; return ErrNotFound }

// Forum threads CRUD
func (s *Store) ListThreads() []models.ForumThread { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.ForumThread(nil), s.threads...) }
func (s *Store) GetThread(id int)(models.ForumThread,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.threads{ if v.ID==id{return v,nil}}; return models.ForumThread{},ErrNotFound }
func (s *Store) CreateThread(v models.ForumThread) models.ForumThread { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.id(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Status==""{v.Status="ouvert"}; s.threads=append(s.threads,v); return v }
func (s *Store) UpdateThread(id int,v models.ForumThread)(models.ForumThread,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.threads{ if s.threads[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.threads[i].CreatedAt}; s.threads[i]=v; return v,nil } }; return models.ForumThread{},ErrNotFound }
func (s *Store) DeleteThread(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.threads{ if v.ID==id{ s.threads=append(s.threads[:i],s.threads[i+1:]...); return nil } }; return ErrNotFound }

// Forum posts CRUD
func (s *Store) ListPosts() []models.ForumPost { s.mu.RLock(); defer s.mu.RUnlock(); return append([]models.ForumPost(nil), s.posts...) }
func (s *Store) GetPost(id int)(models.ForumPost,error){ s.mu.RLock(); defer s.mu.RUnlock(); for _,v:=range s.posts{ if v.ID==id{return v,nil}}; return models.ForumPost{},ErrNotFound }
func (s *Store) CreatePost(v models.ForumPost) models.ForumPost { s.mu.Lock(); defer s.mu.Unlock(); v.ID=s.id(); if v.CreatedAt.IsZero(){v.CreatedAt=time.Now()}; if v.Status==""{v.Status="visible"}; s.posts=append(s.posts,v); return v }
func (s *Store) UpdatePost(id int,v models.ForumPost)(models.ForumPost,error){ s.mu.Lock(); defer s.mu.Unlock(); for i:=range s.posts{ if s.posts[i].ID==id{ v.ID=id; if v.CreatedAt.IsZero(){v.CreatedAt=s.posts[i].CreatedAt}; s.posts[i]=v; return v,nil } }; return models.ForumPost{},ErrNotFound }
func (s *Store) DeletePost(id int) error { s.mu.Lock(); defer s.mu.Unlock(); for i,v:=range s.posts{ if v.ID==id{ s.posts=append(s.posts[:i],s.posts[i+1:]...); return nil } }; return ErrNotFound }
