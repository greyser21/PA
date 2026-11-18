package handlers

import (
	"net/http"
	"strconv"

	"upcycleconnect/api/internal/store"

	"github.com/gin-gonic/gin"
)

func paramID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalide"})
		return 0, false
	}
	return id, true
}

func handleErr(c *gin.Context, err error) bool {
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
