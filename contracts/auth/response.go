package auth

import (
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type RegistrationStatusResponse struct {
	common.BaseErrorResponse
	Status RegistrationStatusDetails `json:"data"`
}

type RegistrationResponse struct {
	common.BaseErrorResponse
	RegistrationID RegistrationID `json:"data"`
}
