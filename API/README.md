# UpcycleConnect API (Go + Gin)

Backend API pour le projet UpcycleConnect (Projet annuel ESGI 2025-2026).

## Lancer
```bash
cd API
go mod tidy
go run ./cmd/server
```
Par défaut: `http://localhost:8080`

## Structure
- `cmd/server` : point d'entrée
- `internal/config` : configuration
- `internal/router` : enregistrement des routes
- `internal/middleware` : CORS, auth (stub/RBAC)
- `internal/models` : modèles de domaine
- `internal/store` : stockage en mémoire (CRUD + seed)
- `internal/handlers` : handlers REST par domaine
- `internal/services` : intégrations (Stripe, PDF, notifications) en stubs extensibles
- `internal/i18n` : base multilingue

## Principales ressources CRUD
`/api/users`, `/api/announcements`, `/api/containers`, `/api/deposits`, `/api/projects`, `/api/events`, `/api/catalog`, `/api/bookings`, `/api/contracts`, `/api/advice`, `/api/notifications` + dashboards `/api/*/dashboard`.
