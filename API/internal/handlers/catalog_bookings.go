package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/models"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterCatalogBookings(rg *gin.RouterGroup, st *store.Store) {
	rc := rg.Group("/catalog")
	rc.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListCatalog()}) })
	rc.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetCatalog(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rc.POST("", func(c *gin.Context) { var v models.CatalogItem; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateCatalog(v)) })
	rc.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.CatalogItem; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateCatalog(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rc.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteCatalog(id)){return}; c.Status(http.StatusNoContent) })

	rb := rg.Group("/bookings")
	rb.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListBookings()}) })
	rb.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetBooking(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rb.POST("", func(c *gin.Context) { var v models.Booking; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateBooking(v)) })
	rb.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Booking; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateBooking(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rb.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteBooking(id)){return}; c.Status(http.StatusNoContent) })
}
