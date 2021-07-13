package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler defines contracts for ping handler.
type Handler interface {
	Ping(c *gin.Context)
}

// handler implements Handler interface.
type handler struct{}

// NewHandler returns a new instance of the handler.
func NewHandler() Handler {
	return &handler{}
}

// Ping will handle ping logic.
func (h *handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
