package handlers

import (
	"net/http"
	"strconv"

	"upcycleconnect/internal/models"
	"upcycleconnect/internal/store"

	"github.com/gin-gonic/gin"
)

type API struct {
	Store *store.Store
}

func New(s *store.Store) *API { return &API{Store: s} }

func (a *API) Register(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", a.health)

		// Dashboards
		api.GET("/user/dashboard", a.userDashboard)
		api.GET("/pro/dashboard", a.proDashboard)
		api.GET("/employee/dashboard", a.employeeDashboard)
		api.GET("/admin/dashboard", a.adminDashboard)

		// Users CRUD
		api.GET("/users", a.listUsers)
		api.GET("/users/:id", a.getUser)
		api.POST("/users", a.createUser)
		api.PUT("/users/:id", a.updateUser)
		api.DELETE("/users/:id", a.deleteUser)

		// Announcements CRUD
		api.GET("/announcements", a.listAnnouncements)
		api.GET("/announcements/:id", a.getAnnouncement)
		api.POST("/announcements", a.createAnnouncement)
		api.PUT("/announcements/:id", a.updateAnnouncement)
		api.DELETE("/announcements/:id", a.deleteAnnouncement)

		// Containers CRUD
		api.GET("/containers", a.listContainers)
		api.GET("/containers/:id", a.getContainer)
		api.POST("/containers", a.createContainer)
		api.PUT("/containers/:id", a.updateContainer)
		api.DELETE("/containers/:id", a.deleteContainer)

		// Deposits CRUD
		api.GET("/deposits", a.listDeposits)
		api.GET("/deposits/:id", a.getDeposit)
		api.POST("/deposits", a.createDeposit)
		api.PUT("/deposits/:id", a.updateDeposit)
		api.DELETE("/deposits/:id", a.deleteDeposit)

		// Projects CRUD
		api.GET("/projects", a.listProjects)
		api.GET("/projects/:id", a.getProject)
		api.POST("/projects", a.createProject)
		api.PUT("/projects/:id", a.updateProject)
		api.DELETE("/projects/:id", a.deleteProject)

		// Events CRUD
		api.GET("/events", a.listEvents)
		api.GET("/events/:id", a.getEvent)
		api.POST("/events", a.createEvent)
		api.PUT("/events/:id", a.updateEvent)
		api.DELETE("/events/:id", a.deleteEvent)

		// Catalog CRUD
		api.GET("/catalog", a.listCatalog)
		api.GET("/catalog/:id", a.getCatalogItem)
		api.POST("/catalog", a.createCatalogItem)
		api.PUT("/catalog/:id", a.updateCatalogItem)
		api.DELETE("/catalog/:id", a.deleteCatalogItem)

		// Bookings CRUD
		api.GET("/bookings", a.listBookings)
		api.GET("/bookings/:id", a.getBooking)
		api.POST("/bookings", a.createBooking)
		api.PUT("/bookings/:id", a.updateBooking)
		api.DELETE("/bookings/:id", a.deleteBooking)

		// Advice CRUD
		api.GET("/advice", a.listAdvice)
		api.GET("/advice/:id", a.getAdvice)
		api.POST("/advice", a.createAdvice)
		api.PUT("/advice/:id", a.updateAdvice)
		api.DELETE("/advice/:id", a.deleteAdvice)

		// Notifications (C,R,D + list)
		api.GET("/notifications", a.listNotifications)
		api.GET("/notifications/:id", a.getNotification)
		api.POST("/notifications", a.createNotification)
		api.DELETE("/notifications/:id", a.deleteNotification)

		// Contracts CRUD
		api.GET("/contracts", a.listContracts)
		api.GET("/contracts/:id", a.getContract)
		api.POST("/contracts", a.createContract)
		api.PUT("/contracts/:id", a.updateContract)
		api.DELETE("/contracts/:id", a.deleteContract)

		// Payments CRUD (Stripe integration point)
		api.GET("/payments", a.listPayments)
		api.GET("/payments/:id", a.getPayment)
		api.POST("/payments", a.createPayment)
		api.PUT("/payments/:id", a.updatePayment)
		api.DELETE("/payments/:id", a.deletePayment)

		// PDF documents
		api.GET("/documents", a.listDocuments)
		api.GET("/documents/:id", a.getDocument)
		api.POST("/documents", a.createDocument)
		api.PUT("/documents/:id", a.updateDocument)
		api.DELETE("/documents/:id", a.deleteDocument)

		// Push subscriptions
		api.GET("/push-subscriptions", a.listPushSubs)
		api.GET("/push-subscriptions/:id", a.getPushSub)
		api.POST("/push-subscriptions", a.createPushSub)
		api.DELETE("/push-subscriptions/:id", a.deletePushSub)

		// Forum
		api.GET("/forum/threads", a.listThreads)
		api.GET("/forum/threads/:id", a.getThread)
		api.POST("/forum/threads", a.createThread)
		api.PUT("/forum/threads/:id", a.updateThread)
		api.DELETE("/forum/threads/:id", a.deleteThread)
		api.GET("/forum/posts", a.listPosts)
		api.GET("/forum/posts/:id", a.getPost)
		api.POST("/forum/posts", a.createPost)
		api.PUT("/forum/posts/:id", a.updatePost)
		api.DELETE("/forum/posts/:id", a.deletePost)
	}
}

func paramID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalide"})
		return 0, false
	}
	return id, true
}

func notFoundOrOK(c *gin.Context, err error) bool {
	if err == store.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "ressource introuvable"})
		return true
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}
	return false
}

func (a *API) health(c *gin.Context)           { c.JSON(http.StatusOK, a.Store.Health()) }
func (a *API) userDashboard(c *gin.Context)    { c.JSON(http.StatusOK, a.Store.UserDashboard()) }
func (a *API) proDashboard(c *gin.Context)     { c.JSON(http.StatusOK, a.Store.ProDashboard()) }
func (a *API) employeeDashboard(c *gin.Context){ c.JSON(http.StatusOK, a.Store.EmployeeDashboard()) }
func (a *API) adminDashboard(c *gin.Context)   { c.JSON(http.StatusOK, a.Store.AdminDashboard()) }

// Users
func (a *API) listUsers(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListUsers()}) }
func (a *API) getUser(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetUser(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createUser(c *gin.Context) {
	var v models.User
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateUser(v))
}
func (a *API) updateUser(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.User
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateUser(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteUser(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteUser(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Announcements
func (a *API) listAnnouncements(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListAnnouncements()}) }
func (a *API) getAnnouncement(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetAnnouncement(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createAnnouncement(c *gin.Context) {
	var v models.Announcement
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateAnnouncement(v))
}
func (a *API) updateAnnouncement(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Announcement
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateAnnouncement(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteAnnouncement(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteAnnouncement(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Containers
func (a *API) listContainers(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListContainers()}) }
func (a *API) getContainer(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetContainer(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createContainer(c *gin.Context) {
	var v models.Container
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateContainer(v))
}
func (a *API) updateContainer(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Container
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateContainer(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteContainer(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteContainer(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Deposits
func (a *API) listDeposits(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListDeposits()}) }
func (a *API) getDeposit(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetDeposit(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createDeposit(c *gin.Context) {
	var v models.Deposit
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateDeposit(v))
}
func (a *API) updateDeposit(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Deposit
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateDeposit(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteDeposit(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteDeposit(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Projects
func (a *API) listProjects(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListProjects()}) }
func (a *API) getProject(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetProject(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createProject(c *gin.Context) {
	var v models.Project
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateProject(v))
}
func (a *API) updateProject(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Project
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateProject(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteProject(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteProject(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Events
func (a *API) listEvents(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListEvents()}) }
func (a *API) getEvent(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetEvent(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createEvent(c *gin.Context) {
	var v models.Event
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateEvent(v))
}
func (a *API) updateEvent(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Event
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateEvent(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteEvent(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteEvent(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Catalog
func (a *API) listCatalog(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListCatalog()}) }
func (a *API) getCatalogItem(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetCatalogItem(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createCatalogItem(c *gin.Context) {
	var v models.CatalogItem
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateCatalogItem(v))
}
func (a *API) updateCatalogItem(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.CatalogItem
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateCatalogItem(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteCatalogItem(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteCatalogItem(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Bookings
func (a *API) listBookings(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListBookings()}) }
func (a *API) getBooking(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetBooking(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createBooking(c *gin.Context) {
	var v models.Booking
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateBooking(v))
}
func (a *API) updateBooking(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Booking
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateBooking(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteBooking(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteBooking(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Advice
func (a *API) listAdvice(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListAdvice()}) }
func (a *API) getAdvice(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetAdvice(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createAdvice(c *gin.Context) {
	var v models.Advice
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateAdvice(v))
}
func (a *API) updateAdvice(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Advice
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateAdvice(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteAdvice(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteAdvice(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Notifications
func (a *API) listNotifications(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListNotifications()}) }
func (a *API) getNotification(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetNotification(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createNotification(c *gin.Context) {
	var v models.Notification
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateNotification(v))
}
func (a *API) deleteNotification(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteNotification(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Contracts
func (a *API) listContracts(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListContracts()}) }
func (a *API) getContract(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	v, err := a.Store.GetContract(id); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, v)
}
func (a *API) createContract(c *gin.Context) {
	var v models.Contract
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	c.JSON(http.StatusCreated, a.Store.CreateContract(v))
}
func (a *API) updateContract(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	var v models.Contract
	if err := c.ShouldBindJSON(&v); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }
	out, err := a.Store.UpdateContract(id, v); if notFoundOrOK(c, err) { return }
	c.JSON(http.StatusOK, out)
}
func (a *API) deleteContract(c *gin.Context) {
	id, ok := paramID(c); if !ok { return }
	if err := a.Store.DeleteContract(id); notFoundOrOK(c, err) { return }
	c.Status(http.StatusNoContent)
}

// Payments
func (a *API) listPayments(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListPayments()}) }
func (a *API) getPayment(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=a.Store.GetPayment(id); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,v) }
func (a *API) createPayment(c *gin.Context) { var v models.Payment; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; c.JSON(http.StatusCreated, a.Store.CreatePayment(v)) }
func (a *API) updatePayment(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.Payment; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; out,err:=a.Store.UpdatePayment(id,v); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,out) }
func (a *API) deletePayment(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if notFoundOrOK(c, a.Store.DeletePayment(id)){return}; c.Status(http.StatusNoContent) }

// Documents
func (a *API) listDocuments(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListDocuments()}) }
func (a *API) getDocument(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=a.Store.GetDocument(id); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,v) }
func (a *API) createDocument(c *gin.Context) { var v models.PDFDocument; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; c.JSON(http.StatusCreated, a.Store.CreateDocument(v)) }
func (a *API) updateDocument(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.PDFDocument; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; out,err:=a.Store.UpdateDocument(id,v); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,out) }
func (a *API) deleteDocument(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if notFoundOrOK(c, a.Store.DeleteDocument(id)){return}; c.Status(http.StatusNoContent) }

// Push subscriptions
func (a *API) listPushSubs(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListPushSubs()}) }
func (a *API) getPushSub(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=a.Store.GetPushSub(id); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,v) }
func (a *API) createPushSub(c *gin.Context) { var v models.PushSubscription; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; c.JSON(http.StatusCreated, a.Store.CreatePushSub(v)) }
func (a *API) deletePushSub(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if notFoundOrOK(c, a.Store.DeletePushSub(id)){return}; c.Status(http.StatusNoContent) }

// Forum
func (a *API) listThreads(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListThreads()}) }
func (a *API) getThread(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=a.Store.GetThread(id); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,v) }
func (a *API) createThread(c *gin.Context) { var v models.ForumThread; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; c.JSON(http.StatusCreated, a.Store.CreateThread(v)) }
func (a *API) updateThread(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.ForumThread; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; out,err:=a.Store.UpdateThread(id,v); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,out) }
func (a *API) deleteThread(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if notFoundOrOK(c, a.Store.DeleteThread(id)){return}; c.Status(http.StatusNoContent) }
func (a *API) listPosts(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"items": a.Store.ListPosts()}) }
func (a *API) getPost(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; v,err:=a.Store.GetPost(id); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,v) }
func (a *API) createPost(c *gin.Context) { var v models.ForumPost; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; c.JSON(http.StatusCreated, a.Store.CreatePost(v)) }
func (a *API) updatePost(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; var v models.ForumPost; if err:=c.ShouldBindJSON(&v); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"données invalides","details":err.Error()}); return }; out,err:=a.Store.UpdatePost(id,v); if notFoundOrOK(c,err){return}; c.JSON(http.StatusOK,out) }
func (a *API) deletePost(c *gin.Context) { id,ok:=paramID(c); if !ok{return}; if notFoundOrOK(c, a.Store.DeletePost(id)){return}; c.Status(http.StatusNoContent) }
