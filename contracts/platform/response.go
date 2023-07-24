package platform

import "github.com/onqlavelabs/onqlave.core/contracts/common"

type StatusResponse struct {
	common.BaseErrorResponse
	Data Status `json:"data"`
}
