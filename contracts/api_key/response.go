package api_key

import (
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type DetailResponse struct {
	common.BaseErrorResponse
	Data APIKey `json:"data"`
}

type ListResponse struct {
	common.BaseErrorResponse
	Data APIKeys `json:"data"`
}

type BaseResponse struct {
	common.BaseErrorResponse
	Data BaseInfo `json:"data"`
}

type SensitiveDataResponse struct {
	common.BaseErrorResponse
	Data SensitiveData `json:"data"`
}
