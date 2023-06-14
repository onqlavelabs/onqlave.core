package requests

import "github.com/onqlavelabs/onqlave.core/contracts/auth"

type RegistrationRequest struct {
	Registration contracts.RegistrationDetails `json:"registration" validate:"required"`
}
