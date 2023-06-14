package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type AddAPIKeyResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKey `json:"data"`
}
