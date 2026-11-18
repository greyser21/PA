package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/models"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterContainersDeposits(rg *gin.RouterGroup, st *store.Store) {
	rc := rg.Group("/containers")
	rc.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListContainers()}) })
	rc.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetContainer(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rc.POST("", func(c *gin.Context) { var v models.Container; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateContainer(v)) })
	rc.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Container; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateContainer(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rc.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteContainer(id)){return}; c.Status(http.StatusNoContent) })

	rd := rg.Group("/deposits")
	rd.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListDeposits()}) })
	rd.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetDeposit(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rd.POST("", func(c *gin.Context) { var v models.Deposit; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateDeposit(v)) })
	rd.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Deposit; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateDeposit(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rd.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteDeposit(id)){return}; c.Status(http.StatusNoContent) })
}
