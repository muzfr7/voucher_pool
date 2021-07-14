package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	customerDomain "github.com/muzfr7/voucher_pool/app/domain/customer"
	"github.com/muzfr7/voucher_pool/app/interfaces"
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
	// decode json body into payload
	var payload = &customerDomain.CreateCustomerRequestDTO{}
	if err := c.ShouldBindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.NewBadRequestError("invalid JSON body"))
		return
	}

	// validate payload data
	if err := payload.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, interfaces.NewBadRequestError(err.Error()))
		return
	}

	// create new customer
	if err := h.customerSVC.Create(payload); err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, nil)
}
