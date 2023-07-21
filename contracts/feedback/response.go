package feedback

import "github.com/onqlavelabs/onqlave.core/contracts/common"

type AddFeedbackResponse struct {
	common.BaseErrorResponse
	Data *NewFeedback `json:"data"`
}
