package requests

import (
	"github.com/onqlavelabs/onqlave.core/contracts/arx"
)

type AddClusterRequest struct {
	Cluster contracts.NewCluster `json:"cluster" validate:"required"`
}
