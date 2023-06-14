package requests

import "github.com/onqlavelabs/onqlave.core/contracts/application"

type AddApplicationRequest struct {
	Application contracts.NewApplication `json:"application" validate:"required"`
}
