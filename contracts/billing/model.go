package billing

import (
	"github.com/onqlavelabs/onqlave.core/contracts/acl"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
	"time"
)

type BillingContract struct {
	ID           string `json:"id" validate:"required,max=150"`
	CompanyName  string `json:"company_name" validate:"required,max=150"`
	BillingEmail string `json:"billing_email" validate:"required,email"`
	UserID       string `json:"user_id" validate:"required"`
}

type Country struct {
	CountryCode string `json:"country_code,omitempty" validate:"required,max=20"`
	CountryName string `json:"country_name" validate:"required,max=50"`
	Icon        string `json:"icon" validate:"required,max=50"`
}

type ReferenceInformation struct {
	Countries []Country `json:"countries"`
}

type Plan struct {
	Type          string     `json:"type,omitempty"`
	Name          string     `json:"name,omitempty"`
	Description   string     `json:"description,omitempty"`
	EndDate       *time.Time `json:"end_date,omitempty"`
	RemainingDays uint       `json:"remaining_days,omitempty"`
}

type TaxSetting struct {
	TaxID      string `json:"tax_id"`
	VATNumber  string `json:"vat_number"`
	VATEnabled bool   `json:"vat_enabled"`
}

type Address struct {
	CountryCode  string `json:"country_code"`
	AddressLine1 string `json:"address_line1"`
	ZipCode      string `json:"zip_code"`
	State        string `json:"state"`
}

type BillingDetail struct {
	CompanyName  string     `json:"company_name"`
	BillingEmail string     `json:"billing_email"`
	TaxSetting   TaxSetting `json:"tax_settings"`
	PhoneNumber  string     `json:"phone_number"`
	Address      Address    `json:"address"`
}

type PaymentMethod struct {
	ID             string `json:"id"`
	CardNumber     string `json:"card_number"`
	NameOnCard     string `json:"name_on_card"`
	ExpiryYear     uint64 `json:"expiry_year"`
	ExpiryMonth    uint64 `json:"expiry_month"`
	IsDefault      bool   `json:"is_default"`
	Brand          string `json:"brand"`
	CardHolderName string `json:"card_holder_name"`
}

type BillingEvent struct {
	ID               string `json:"id"`
	AccountID        string `json:"account_id"`
	UserID           string `json:"user_id"`
	EventType        string `json:"event_type"`
	EventDescription string `json:"event_description"`
}

type BillingSubscriptionType struct {
	Type         string                           `json:"type"`
	Name         string                           `json:"name"`
	Description  string                           `json:"description"`
	CanUpgradeTo string                           `json:"can_upgrade_to,omitempty"`
	IsPopular    bool                             `json:"is_popular"`
	Order        uint64                           `json:"order"`
	IsEnterprise bool                             `json:"is_enterprise"`
	Products     []BillingSubscriptionTypeProduct `json:"products"`
}

type BillingSubscriptionTypeProduct struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Order    uint64 `json:"order"`
	IsMain   bool   `json:"is_main"`
	Type     string `json:"type"`
	Billable bool   `json:"billable"`
	FreeTier int64  `json:"free_tier"`
	MaxTier  int64  `json:"max_tier"`
}

type BillingFAQ struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDefault   bool   `json:"is_default"`
	Icon        string `json:"icon,omitempty"`
}

type BillingAccountInfo struct {
	ACL               acl.ACL                   `json:"acl"`
	BillingDetails    BillingDetail             `json:"billing_details"`
	Model             ReferenceInformation      `json:"model"`
	CurrentPlan       *Plan                     `json:"current_plan"`
	UpgradePlan       *Plan                     `json:"upgrade_plan"`
	PaymentMethods    []PaymentMethod           `json:"payment_methods"`
	SubscriptionPlans []BillingSubscriptionType `json:"subscription_plans"`
	FAQs              []BillingFAQ              `json:"faqs"`
	Events            []BillingEvent            `json:"events"`
}

type GetBillingAccountResponse struct {
	common.BaseErrorResponse
	BillingAccount BillingAccountInfo `json:"data"`
}

type GetListCountriesResponse struct {
	common.BaseErrorResponse
	Countries []Country `json:"countries"`
}

type UpgradeSubscriptionRequest struct {
	SubscriptionType string `json:"subscription_type" validate:"required,max=100"`
}

type UpgradeSubscriptionResponse struct {
	common.BaseErrorResponse
	UpgradeSubscription Subscription `json:"data"`
}

type Subscription struct {
	AccountID        string     `json:"account_id"`
	StripeCustomerID string     `json:"stripe_customer_id,omitempty"`
	SubscriptionType string     `json:"subscription_type"`
	SubscriptionID   string     `json:"subscription_id,omitempty"`
	CycleEndDate     *time.Time `json:"cycle_end_date,omitempty"`
	CycleStartDate   *time.Time `json:"cycle_start_date,omitempty"`
	IsActive         bool       `json:"is_active"`
}

type CancelSubscriptionResponse struct {
	common.BaseErrorResponse
	CancelSubscription Subscription `json:"data"`
}

type BillingAccount struct {
	CompanyName     string  `json:"company_name" validate:"required,max=150"`
	CompanyUri      string  `json:"company_uri" validate:"omitempty,url"`
	AddressLine1    string  `json:"address_line1" validate:""`
	AddressLine2    string  `json:"address_line2" validate:""`
	City            string  `json:"city" validate:""`
	ZipCode         string  `json:"zip_code" validate:"required_with=AddressLine1 AddressLine2"`
	State           string  `json:"state" validate:"required_with=AddressLine1 AddressLine2"`
	CountryCode     string  `json:"country_code" validate:"required,max=20"`
	BillingEmail    string  `json:"billing_email" validate:"required,email"`
	TaxID           string  `json:"tax_id" validate:"omitempty,max=100"`
	VatNumber       string  `json:"vat_number" validate:"omitempty,max=150"`
	TaxJurisdiction string  `json:"tax_jurisdiction" validate:"omitempty,max=100"`
	TaxPercentage   float64 `json:"tax_percentage" validate:"omitempty,min=0,max=100"`
	PhoneNumber     string  `json:"phone_number" validate:"omitempty,e164"`
}

type BillingAccountRequest struct {
	BillingAccount
}

type UpdateBillingAccountResponse struct {
	common.BaseErrorResponse
	BillingAccountID string `json:"data"`
}

type CurrentSubscription struct {
	AccountID        string `json:"account_id"`
	SubscriptionType string `json:"subscription_type"`
}

type PaymentToken struct {
	PaymentToken string `json:"payment_token" validate:"required,max=150"`
	IsDefault    *bool  `json:"is_default" validate:"required"`
}

type SubscriptionTypeProduct struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Billable bool   `json:"billable"`
	MaxTier  int64  `json:"max_tier"`
}
