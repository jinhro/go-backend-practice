package api

import (
	db "jinhro/go-backend-practice/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for the service
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new Server instance
func NewServer(store *db.Store) *Server {	
	server := &Server {
		store: store,
	}
	router := gin.Default();
	
	//define routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router;
	return server
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
