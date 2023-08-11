package auth

type RegistrationStatus struct {
	Request RegistrationToken `json:"request" validate:"required"`
}

type RegistrationRequest struct {
	Registration RegistrationDetails `json:"registration" validate:"required"`
}

type RegistrationValidation struct {
	Validation RegistrationToken `json:"validation" validate:"required"`
}

type LoginRequest struct {
	LoginDetails LoginDetails `json:"login" validate:"required"`
}
