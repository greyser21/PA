package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/models"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterUsers(rg *gin.RouterGroup, st *store.Store) {
	r := rg.Group("/users")
	r.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListUsers()}) })
	r.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetUser(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	r.POST("", func(c *gin.Context) { var v models.User; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateUser(v)) })
	r.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.User; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateUser(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	r.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteUser(id)){return}; c.Status(http.StatusNoContent) })
}
