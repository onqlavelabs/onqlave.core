package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/application"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type SealApplicationResponse struct {
	common.BaseErrorResponse
	Data contracts.ApplicationStatus `json:"data"`
}
