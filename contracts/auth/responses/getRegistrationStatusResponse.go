package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/auth"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type RegistrationStatusResponse struct {
	common.BaseErrorResponse
	Status contracts.RegistrationStatusDetails `json:"data"`
}
