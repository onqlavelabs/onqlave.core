package enumerations

type ApiKeyStatus string

const (
	ApiKeyStatusPending  ApiKeyStatus = "pending"
	ApiKeyStatusActive   ApiKeyStatus = "active"
	ApiKeyStatusDisabled ApiKeyStatus = "disabled"
	ApiKeyStatusDeleted  ApiKeyStatus = "deleted"
	ApiKeyStatusFailed   ApiKeyStatus = "failed"
)

func (status ApiKeyStatus) String() string {
	return string(status)
}
