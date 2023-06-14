package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type UpdateClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingClusterWithDetails `json:"data"`
}
