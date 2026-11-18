# Documentation Backend Go — UpcycleConnect

Date : 2026-11-12  
Projet : UpcycleConnect — Projet annuel 2 ESGI (2025-2026)

## 1. Périmètre réalisé (backend/Go)
Conformément à la demande (« tout le backend/go »), le travail s’est concentré sur le **backend en Go (Gin)** et l’**API REST** :
- Extraction des consignes depuis `documents/Sujet du projet.pdf` via `pdftotext` (le PDF ne peut pas être lu directement par le modèle) et prise en compte de `documents/Cahier_des_Charges.md`.
- Backend initial à la racine (`cmd/server`, `internal/*`) + frontend `web/` rendu dynamique.
- **Backend principal structuré dans `API/`** (~27 fichiers, créé avec ~80 s d’attente entre chaque création) : architecture modulaire (config, router, middleware, i18n, models, store, handlers, services).
- **CRUD/API étendus** pour les entités métier nécessaires aux espaces particuliers, pros/artisans, salariés et administration.

> Les parties hors backend Go (Mission 2 infra/réseau/sécurité, Mission 3 modules non-Go, livrables non-code) ne sont pas traitées ici, comme demandé.

## 2. Architecture — Backend principal (`API/`)

```
API/
├── go.mod, go.sum, README.md
├── cmd/server/main.go                 # Entrée serveur Gin
└── internal/
    ├── config/config.go               # Configuration (env: API_ADDR, GIN_MODE, STRIPE_KEY, ONESIGNAL_*, DEFAULT_LANG)
    ├── router/router.go               # Montage routes/middlewares
    ├── middleware/cors.go             # CORS
    ├── middleware/auth.go             # Auth/RBAC (stub extensible)
    ├── i18n/i18n.go                   # Multilingue (X-Lang / ?lang), extensible
    ├── models/
    │   ├── user.go
    │   ├── announcement.go
    │   ├── container.go               # Container + Deposit (codes ouverture/code-barres)
    │   ├── project.go                 # Projets d'upcycling, impact score
    │   ├── event_catalog.go           # Events/formations/ateliers + catalogue
    │   ├── commerce.go                # Bookings, contrats, paiements
    │   └── content.go                 # Conseils, notifications, documents PDF, push subs, forum
    ├── store/store.go                 # Store en mémoire thread-safe, seed, CRUD, dashboards
    ├── handlers/
    │   ├── helpers.go                 # Parsing ID, erreurs 400/404/500
    │   ├── health_dash.go             # /health + dashboards 4 espaces
    │   ├── users.go                   # CRUD utilisateurs
    │   ├── announcements.go           # CRUD annonces
    │   ├── containers_deposits.go     # CRUD conteneurs & dépôts
    │   ├── projects_events.go         # CRUD projets & événements
    │   ├── catalog_bookings.go        # CRUD catalogue & réservations
    │   └── content_admin.go           # CRUD conseils, notifications, contrats, paiements, documents, push, forum
    └── services/
        ├── stripe.go                  # Intégration paiement Stripe (stub prêt à brancher)
        ├── pdf.go                     # Génération/archivage PDF (stub prêt à brancher)
        └── notify.go                  # Notifications push ex. OneSignal (stub prêt à brancher)
```

**Caractéristiques :**
- Go + **Gin**, API REST JSON.
- **CRUD complet** (GET liste, GET/:id, POST, PUT/:id, DELETE/:id selon ressource) sur les entités clés.
- Store en mémoire thread-safe avec données seed, IDs auto, statuts/horodatages/codes (dépôt: `openCode`, `barcode`).
- Endpoints dashboards pour les 4 espaces (particulier, pro/artisan, salarié, admin).
- CORS, base i18n, point d’extension auth/RBAC, services stubs Stripe/PDF/push.

## 3. Endpoints API (`API/`)

Base : `http://localhost:8080/api`

### Santé & dashboards
- `GET /health`
- `GET /user/dashboard`
- `GET /pro/dashboard`
- `GET /employee/dashboard`
- `GET /admin/dashboard`

### CRUD ressources
| Ressource | Endpoints |
|---|---|
| Utilisateurs | `/users`, `/users/:id` |
| Annonces (don/vente, validation) | `/announcements`, `/announcements/:id` |
| Conteneurs/box | `/containers`, `/containers/:id` |
| Dépôts (demandes, codes, récupération) | `/deposits`, `/deposits/:id` |
| Projets d'upcycling | `/projects`, `/projects/:id` |
| Événements/formations/ateliers | `/events`, `/events/:id` |
| Catalogue (services/formations/événements) | `/catalog`, `/catalog/:id` |
| Réservations/planning | `/bookings`, `/bookings/:id` |
| Conseils/news | `/advice`, `/advice/:id` |
| Notifications | `/notifications`, `/notifications/:id` (`GET/POST/DELETE`) |
| Contrats/abonnements/publicité | `/contracts`, `/contracts/:id` |
| Paiements (Stripe) | `/payments`, `/payments/:id` |
| Documents PDF | `/documents`, `/documents/:id` |
| Abonnements push | `/push-subscriptions`, `/push-subscriptions/:id` (`GET/POST/DELETE`) |
| Forum — fils | `/forum/threads`, `/forum/threads/:id` |
| Forum — messages/modération | `/forum/posts`, `/forum/posts/:id` |

**Conventions** : listes `{ "items": [...] }`, création `201`, suppression `204`, `400` données invalides, `404` introuvable.

## 4. Lancement et vérification

### Backend principal (`API/`)
```bash
cd API
go mod tidy
go build ./cmd/server
go run ./cmd/server
# API: http://localhost:8080/api/health
```

### Backend racine (prototype initial, toujours fonctionnel)
```bash
go mod tidy
go build ./cmd/server
go run ./cmd/server
# Sert web/ + API sur http://localhost:8080
```

**Vérifications effectuées** : `API/` → `go mod tidy`, `go build ./cmd/server` OK, `/api/health` répond `{"service":"upcycleconnect-api","status":"ok"}`. Racine → build OK, API et pages web dynamiques opérationnelles.

## 5. Fichiers créés/modifiés (récapitulatif)

### Backend racine (modifié/complété)
- `go.mod`/`go.sum` : dépendance Gin.
- `cmd/server/main.go` : serveur Gin, API + statique `web/`.
- `internal/models/models.go` : modèles domaine + dashboard + paiements/docs/push/forum.
- `internal/store/store.go` : store mémoire, seed, CRUD étendu (incl. paiements/docs/push/forum), dashboards.
- `internal/handlers/api.go` : handlers/routes REST CRUD complets + dashboards.
- `go.md` : documentation.

### Frontend (connecté à l’API, placeholders remplis)
- `web/user.html`, `web/pro.html`, `web/employee.html`, `web/admin.html` : `data-page`, zones dynamiques, formulaire annonce/dépôt.
- `web/js/app.js` : fetch API, rendu stats/listes, actions/formulaires.
- `web/css/style.css` : styles formulaire/toasts.

### Dossier `API/` (créé, backend principal)
- 27 fichiers listés en section 2 (config, router, middleware, i18n, 7 fichiers models, store, 7 fichiers handlers, 3 services, cmd, go.mod, README). Création séquentielle avec ~80 s entre chaque fichier (`/tmp/create_api.log`).

## 6. Couverture fonctionnelle backend vs consignes Mission 1

**Implémenté (backend Go/API)** :
- API Go (Gin) avec endpoints pour traitements principaux, CRUD des entités : utilisateurs/acteurs, annonces, conteneurs/dépôts (avec génération codes ouverture/code-barres), projets, événements/formations/ateliers, catalogue, réservations/planning, conseils, notifications, contrats/abonnements/publicité, paiements, documents PDF (métadonnées/chemins), abonnements push, forum (threads/posts/modération via statut).
- Dashboards pour espaces particulier, pro/artisan, salarié, admin.
- Structure prête pour : Stripe (`services/stripe`, `/payments`), génération/archivage PDF (`services/pdf`, `/documents`), notifications push (`services/notify`, `/push-subscriptions`), multilingue (`i18n`), auth/RBAC (`middleware/auth`).

**À compléter pour une mise en production conforme totale** :
- Persistance en base de données (PostgreSQL/MySQL/SQLite) + migrations (actuellement mémoire).
- Authentification/autorisation réelle (JWT/sessions), contrôle d’accès strict par rôle.
- Intégration effective Stripe (clés, webhooks, Checkout/PaymentIntent) et génération PDF réelle (librairie + stockage/accès doublon).
- Envoi push réel (OneSignal/API) et gestion avancée des notifications.
- i18n complet via fichiers de ressources/langues chargeables sans modification de code.
- Workflows métier avancés détaillés (validation admin fine multi-étapes, service Check complet, calcul Upcycling Score, gestion financière détaillée, planning riche, forums complets, etc.), validations métier renforcées, et UI back/front exhaustive sur tous les CRUD.
- Déploiement serveur (réécriture d’URL, codes d’erreur), Docker/compose en prod, configuration Cloud selon validation.

Le backend Go/API fourni constitue une base structurée et fonctionnelle (CRUD + endpoints + architecture) couvrant les domaines demandés côté applicatif, à étendre avec les intégrations et la persistance pour la conformité production complète.