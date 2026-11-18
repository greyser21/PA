package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterHealth(rg *gin.RouterGroup, st *store.Store) {
	rg.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, st.Health()) })
}

func RegisterDashboards(rg *gin.RouterGroup, st *store.Store) {
	rg.GET("/user/dashboard", func(c *gin.Context) { c.JSON(http.StatusOK, st.DashUser()) })
	rg.GET("/pro/dashboard", func(c *gin.Context) { c.JSON(http.StatusOK, st.DashPro()) })
	rg.GET("/employee/dashboard", func(c *gin.Context) { c.JSON(http.StatusOK, st.DashEmployee()) })
	rg.GET("/admin/dashboard", func(c *gin.Context) { c.JSON(http.StatusOK, st.DashAdmin()) })
}
