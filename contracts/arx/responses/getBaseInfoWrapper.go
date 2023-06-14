package responses

import (
	"github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type GetClusterBaseInfoWrapper struct {
	common.BaseErrorResponse
	Data contracts.ArxModelWrapper `json:"data"`
}
