package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/models"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterAnnouncements(rg *gin.RouterGroup, st *store.Store) {
	r := rg.Group("/announcements")
	r.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListAnnouncements()}) })
	r.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetAnnouncement(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	r.POST("", func(c *gin.Context) { var v models.Announcement; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateAnnouncement(v)) })
	r.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Announcement; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateAnnouncement(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	r.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteAnnouncement(id)){return}; c.Status(http.StatusNoContent) })
}
