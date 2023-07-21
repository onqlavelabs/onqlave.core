package billing

type CreatePaymentMethod struct {
	Request PaymentToken `json:"request" validate:"required"`
}

type UpdatePaymentMethod struct {
	Request PaymentToken `json:"request" validate:"required"`
}
