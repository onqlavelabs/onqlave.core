package contracts

import "github.com/onqlavelabs/onqlave.core/contracts/acl"

type GetApplicationsResponseWrapper struct {
	ACL          contracts.ACL                    `json:"acl"`
	Applications []ExistingApplicationWithDetails `json:"applications"`
	Models       ApplicationModelWrapper          `json:"model"`
	Statistics   ApplicationStatistics            `json:"statistics"`
}
