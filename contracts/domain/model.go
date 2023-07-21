package domain

import "github.com/onqlavelabs/onqlave.core/contracts/common"

type Domain struct {
	TenantID common.TenantId
	UserID   common.UserId
}
