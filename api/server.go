package api

import (
	db "github.com/arthurlch/job_board_api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store 
	router *gin.Engine
}