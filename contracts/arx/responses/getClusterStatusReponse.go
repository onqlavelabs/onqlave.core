package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type GetClusterStateResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
