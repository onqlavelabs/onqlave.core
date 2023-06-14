package requests

import "github.com/onqlavelabs/onqlave.core/contracts/auth"

type RegistrationStatusRequest struct {
	Request contracts.RegistrationToken `json:"request" validate:"required"`
}
