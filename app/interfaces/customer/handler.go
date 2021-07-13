package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"
)

// handler implements customerDomain.Handler interface.
type handler struct {
	customerSVC customerDomain.Service
}

// NewHandler returns a new instance of the handler.
func NewHandler(s customerDomain.Service) customerDomain.Handler {
	return &handler{
		customerSVC: s,
	}
}

// Create will handle customer creation.
func (h *handler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}
