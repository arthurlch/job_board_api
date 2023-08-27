package api

import (
	db "github.com/arthurlch/job_board_api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/applications", server.createApplication)
	router.GET("/applications", server.listApplications)
	router.GET("/applications/:id", server.getApplicationByID)
	router.PUT("/applications/:id", server.updateApplication)
	router.DELETE("/applications/:id", server.deleteApplication)

	server.router = router 
	return server
}

func (server *Server) Start(address string) error { 
	return server.router.Run(address) 
}

func ErrorResponse(err error) gin.H { 
	return gin.H{"error": err.Error()} 
}
