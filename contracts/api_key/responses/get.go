package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/api_key"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type GetAPIKeysResponse struct {
	common.BaseErrorResponse
	Data contracts.GetAPIKeysResponseWrapper `json:"data"`
}

type GetAPIKeyBaseInformationResponse struct {
	common.BaseErrorResponse
	Data contracts.GetAPIKeyBaseResponse `json:"data"`
}
