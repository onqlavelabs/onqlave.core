package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/auth"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type RegistrationResponse struct {
	common.BaseErrorResponse
	RegistrationID contracts.RegistrationID `json:"data"`
}
