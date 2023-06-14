package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/application"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type GetApplicationsResponse struct {
	common.BaseErrorResponse
	Data contracts.GetApplicationsResponseWrapper `json:"data"`
}
