package handlers

import (
	"net/http"

	"upcycleconnect/api/internal/models"
	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func RegisterContentAdmin(rg *gin.RouterGroup, st *store.Store) {
	ra := rg.Group("/advice")
	ra.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListAdvice()}) })
	ra.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetAdvice(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	ra.POST("", func(c *gin.Context) { var v models.Advice; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateAdvice(v)) })
	ra.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Advice; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateAdvice(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	ra.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteAdvice(id)){return}; c.Status(http.StatusNoContent) })

	rn := rg.Group("/notifications")
	rn.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListNotifications()}) })
	rn.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetNotification(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rn.POST("", func(c *gin.Context) { var v models.Notification; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateNotification(v)) })
	rn.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteNotification(id)){return}; c.Status(http.StatusNoContent) })

	rct := rg.Group("/contracts")
	rct.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListContracts()}) })
	rct.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetContract(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rct.POST("", func(c *gin.Context) { var v models.Contract; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateContract(v)) })
	rct.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Contract; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateContract(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rct.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteContract(id)){return}; c.Status(http.StatusNoContent) })

	rp := rg.Group("/payments")
	rp.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListPayments()}) })
	rp.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetPayment(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rp.POST("", func(c *gin.Context) { var v models.Payment; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreatePayment(v)) })
	rp.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Payment; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdatePayment(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rp.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeletePayment(id)){return}; c.Status(http.StatusNoContent) })

	rd := rg.Group("/documents")
	rd.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListDocuments()}) })
	rd.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetDocument(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rd.POST("", func(c *gin.Context) { var v models.PDFDocument; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateDocument(v)) })
	rd.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.PDFDocument; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateDocument(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	rd.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteDocument(id)){return}; c.Status(http.StatusNoContent) })

	rps := rg.Group("/push-subscriptions")
	rps.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListPushSubs()}) })
	rps.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetPushSub(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	rps.POST("", func(c *gin.Context) { var v models.PushSubscription; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreatePushSub(v)) })
	rps.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeletePushSub(id)){return}; c.Status(http.StatusNoContent) })

	ft := rg.Group("/forum/threads")
	ft.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListThreads()}) })
	ft.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetThread(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	ft.POST("", func(c *gin.Context) { var v models.ForumThread; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreateThread(v)) })
	ft.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.ForumThread; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdateThread(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	ft.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeleteThread(id)){return}; c.Status(http.StatusNoContent) })

	fp := rg.Group("/forum/posts")
	fp.GET("", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": st.ListPosts()}) })
	fp.GET("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=st.GetPost(id); if handleErr(c,err){return}; c.JSON(http.StatusOK,v) })
	fp.POST("", func(c *gin.Context) { var v models.ForumPost; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; c.JSON(http.StatusCreated, st.CreatePost(v)) })
	fp.PUT("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.ForumPost; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides"}); return }; out,err:=st.UpdatePost(id,v); if handleErr(c,err){return}; c.JSON(http.StatusOK,out) })
	fp.DELETE("/:id", func(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if handleErr(c, st.DeletePost(id)){return}; c.Status(http.StatusNoContent) })
}
