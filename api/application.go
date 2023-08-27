package api

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/arthurlch/job_board_api/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) createApplication(c *gin.Context) {
	var req db.InsertApplicationParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := server.store.InsertApplication(context.Background(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "inserted"})
}

func (server *Server) listApplications(c *gin.Context) {
	data, err := server.store.SelectAllApplications(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (server *Server) getApplicationByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := server.store.SelectApplicationByID(context.Background(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (server *Server) updateApplication(c *gin.Context) {
	var req db.UpdateApplicationParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	req.ID = int32(id)
	if err := server.store.UpdateApplication(context.Background(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func (server *Server) deleteApplication(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := server.store.DeleteApplication(context.Background(), int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
} 

