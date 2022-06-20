package domainevents

type DomainEvent interface {
	Topic() string
}

type RegistrationRequestReceived struct {
	EmailAddress     string `json:"email" validate:"email,required,min=4,max=50"`
	Token            string `json:"token" validate:"required,min=4,max=100"`
	LinkValidityTime int32  `json:"validity" validate:"min=15,max=120"`
	Link             string `json:"link" validate:"required,min=10,max=500"`
	Operation        string `json:"operation" validate:"required,min=10,max=50"`
	TenantName       string `json:"tenant_name" validate:"required,min=4,max=100"`
}

func (d *RegistrationRequestReceived) Topic() string {
	return "RegistrationRequestReceived"
}

type RegistrationRequestExpired struct {
	Token string `json:"token" validate:"required,min=4,max=100"`
}

func (d *RegistrationRequestExpired) Topic() string {
	return "RegistrationRequestExpired"
}

type RegistrationEmailSent struct {
	Token   string `json:"token" validate:"required,min=4,max=100"`
	Message string `json:"message" validate:"required"`
	IsError bool   `json:"is_error"`
}

type TenantAdded struct {
	Name      string `json:"tenant_name" validate:"required,min=4,max=100"`
	Id        string `json:"tenant_id" validate:"required,min=4,max=100"`
	Message   string `json:"message" validate:"required"`
	IsError   bool   `json:"is_error"`
	RequestId string `json:"request_id" validate:"required,min=4,max=100"`
}

func (d *TenantAdded) Topic() string {
	return "TenantAdded"
}

func (d *RegistrationEmailSent) Topic() string {
	return "RegistrationEmailSent"
}

type RegistrationRequestCompleted struct {
	EmailAddress string `json:"email" validate:"email,required,min=4,max=50"`
	Token        string `json:"token" validate:"required,min=4,max=100"`
}

func (d *RegistrationRequestCompleted) Topic() string {
	return "RegistrationRequestCompleted"
}
