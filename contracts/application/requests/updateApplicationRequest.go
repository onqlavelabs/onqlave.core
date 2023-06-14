package requests

import "github.com/onqlavelabs/onqlave.core/contracts/application"

type UpdateApplicationRequest struct {
	Application contracts.UpdateApplication `json:"application" validate:"required"`
}
