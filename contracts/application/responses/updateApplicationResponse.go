package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/application"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type UpdateApplicationResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingApplicationWithDetails `json:"data"`
}
