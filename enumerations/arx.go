package enumerations

type ArxStatus string

const (
	ArxStatusActive      ArxStatus = "active"
	ArxStatusSealed      ArxStatus = "sealed"
	ArxStatusUnsealed    ArxStatus = "unsealed"
	ArxStatusInactive    ArxStatus = "inactive"
	ArxStatusPending     ArxStatus = "pending"
	ArxStatusFailed      ArxStatus = "failed"
	ArxStatusDeleted     ArxStatus = "deleted"
	ArxStatusInitiated   ArxStatus = "initiated"
	ArxStatusReInitiated ArxStatus = "reinitiated"
)

func (status ArxStatus) String() string {
	return string(status)
}
