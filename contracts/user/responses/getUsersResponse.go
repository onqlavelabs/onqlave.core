package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type GetUsersResponse struct {
	common.BaseErrorResponse
	Data contracts.GetUsersResponse `json:"data"`
}
