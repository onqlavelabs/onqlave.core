package enumerations

// RegistrationOperation defines the action of the registration record
type RegistrationOperation string

const (
	RegistrationOperationLogin         RegistrationOperation = "login"
	RegistrationOperationSignup        RegistrationOperation = "signup"
	RegistrationOperationResetPassword RegistrationOperation = "reset_password"
	RegistrationOperationInviteUser    RegistrationOperation = "invite_user"
)

func (operation RegistrationOperation) String() string {
	return string(operation)
}

// RegistrationStatus defines the status of the registration record
type RegistrationStatus string

const (
	RegistrationStatusInitiated RegistrationStatus = "initiated"
	RegistrationStatusPending   RegistrationStatus = "pending"
	RegistrationStatusCompleted RegistrationStatus = "completed"
	RegistrationStatusExpired   RegistrationStatus = "expired"
)

func (status RegistrationStatus) String() string {
	return string(status)
}
