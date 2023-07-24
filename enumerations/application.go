package enumerations

type ApplicationStatus string

const (
	ApplicationStatusActive   ApplicationStatus = "active"
	ApplicationStatusArchived ApplicationStatus = "archived"
	ApplicationStatusDisabled ApplicationStatus = "disabled"
	ApplicationStatusPending  ApplicationStatus = "pending"
	ApplicationStatusDeleted  ApplicationStatus = "deleted"
)

func (status ApplicationStatus) String() string {
	return string(status)
}
