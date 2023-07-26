package common

import (
	"regexp"
)

type ArxId string
type ApplicationId string
type TenantId string
type UserId string
type UserEmail string
type RegistrationID string
type DomainId string

func (id *ArxId) Valid() bool {
	return true
}

func (id *ApplicationId) Valid() bool {
	return true
}

func (id *TenantId) Valid() bool {
	pattern := regexp.MustCompile(`--`)
	result := pattern.Split(string(*id), -1)
	return len(result) == 2
}

func EmptyRegistrationID() RegistrationID {
	return RegistrationID("")
}

func EmptyApplicationId() ApplicationId {
	return ApplicationId("")
}

func EmptyUserId() UserId {
	return UserId("")
}

func EmptyTenantId() TenantId {
	return TenantId("")
}

func EmptyArxId() ArxId {
	return ArxId("")
}

func ExisingTenantId(id string) TenantId {
	return TenantId(id)
}

func ExisingUserId(id string) UserId {
	return UserId(id)
}

func ExisingArxId(id string) ArxId {
	return ArxId(id)
}
