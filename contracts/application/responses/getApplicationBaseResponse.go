package responses

import "github.com/onqlavelabs/onqlave.core/contracts/application"

type GetApplicationBaseResponse struct {
	Data contracts.ApplicationModelWrapper `json:"data"`
}
