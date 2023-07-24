package dashboard

import "github.com/onqlavelabs/onqlave.core/contracts/common"

type GetOperationDashboardResponse struct {
	common.BaseErrorResponse
	Dashboard Operations `json:"data"`
}
