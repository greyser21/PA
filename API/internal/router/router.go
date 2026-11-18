package router

import (
	"upcycleconnect/api/internal/config"
	"upcycleconnect/api/internal/handlers"
	"upcycleconnect/api/internal/i18n"
	"upcycleconnect/api/internal/middleware"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, st *store.Store, cfg config.Config) {
	r.Use(middleware.CORS())
	r.Use(i18n.Middleware(cfg.DefaultLang))
	r.Use(middleware.OptionalAuth())

	api := r.Group("/api")
	{
		handlers.RegisterHealth(api, st)
		handlers.RegisterDashboards(api, st)
		handlers.RegisterUsers(api, st)
		handlers.RegisterAnnouncements(api, st)
		handlers.RegisterContainersDeposits(api, st)
		handlers.RegisterProjectsEvents(api, st)
		handlers.RegisterCatalogBookings(api, st)
		handlers.RegisterContentAdmin(api, st)
	}
}
