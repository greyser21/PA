# UpcycleConnect

Projet de fin d'année ESGI 2026 — Plateforme d'upcycling intelligent.

## Structure

```
.
├── cmd/server/      # Entry point du serveur Go (http.ListenAndServe)
├── internal/        # Packages Go internes (handlers, models, etc.)
├── web/             # Frontend (HTML, CSS, JS)
│   ├── css/
│   └── js/
└── go.mod
```

## Lancer le serveur

```sh
go run ./cmd/server
```

Puis ouvrir http://localhost:8080