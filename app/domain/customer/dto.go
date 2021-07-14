package customer

import "errors"

// CreateCustomerRequestDTO is the request payload to POST /customers endpoint.
type CreateCustomerRequestDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate validates required fields in CreateCustomerRequestDTO.
func (dto *CreateCustomerRequestDTO) Validate() error {
	if dto.Name == "" {
		return errors.New("name cannot be blank")
	}

	if dto.Email == "" {
		return errors.New("email cannot be blank")
	}

	return nil
}
