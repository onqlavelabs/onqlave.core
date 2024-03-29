package contracts

import (
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type DetailResponse struct {
	common.BaseErrorResponse
	Data ExistingWithDetail `json:"data"`
}

type StatusResponse struct {
	common.BaseErrorResponse
	Data Status `json:"data"`
}

type ListResponse struct {
	common.BaseErrorResponse
	Data GetListResponse `json:"data"`
}

type BaseInfoResponse struct {
	common.BaseErrorResponse
	Data BaseInfo `json:"data"`
}

type DefaultResponse struct {
	common.BaseErrorResponse
	Data GetDefaultArxWrapper `json:"data"`
}
