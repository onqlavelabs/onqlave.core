package requests

import (
	"github.com/onqlavelabs/onqlave.core/contracts/arx"
)

type UpdateClusterRequest struct {
	Cluster contracts.UpdateCluster `json:"cluster" validate:"required"`
}
