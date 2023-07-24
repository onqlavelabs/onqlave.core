package subscription

type CurrentSubscription struct {
	TenantID         string `json:"tenant_id"`
	SubscriptionType string `json:"subscription_type"`
}
