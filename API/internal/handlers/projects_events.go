package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/models"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterProjectsEvents(rg *gin.RouterGroup, st *store.Store) {
	rp := rg.Group("/projects")
	rp.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListProjects()}) })
	rp.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetProject(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rp.POST("", func(c *gin.Context) { var v models.Project; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateProject(v)) })
	rp.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Project; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateProject(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rp.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteProject(id)){return}; c.Status(http.StatusNoContent) })

	re := rg.Group("/events")
	re.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListEvents()}) })
	re.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetEvent(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	re.POST("", func(c *gin.Context) { var v models.Event; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateEvent(v)) })
	re.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Event; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateEvent(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	re.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteEvent(id)){return}; c.Status(http.StatusNoContent) })
}
