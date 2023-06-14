package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type SealClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ClusterStatus `json:"data"`
}
