package requests

import "github.com/onqlavelabs/onqlave.core/contracts/auth"

type LoginRequest struct {
	LoginDetails contracts.LoginDetails `json:"login" validate:"required"`
}
