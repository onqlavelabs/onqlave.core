package billing

import "github.com/onqlavelabs/onqlave.core/contracts/common"

type GetCurrentSubscription struct {
	common.BaseErrorResponse
	Data CurrentSubscription `json:"data"`
}

type GetSubscriptionTypeProductResponse struct {
	common.BaseErrorResponse
	Data []SubscriptionTypeProduct `json:"data"`
}
