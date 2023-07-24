package feedback

import "time"

type NewFeedback struct {
	ID        string     `json:"id"`
	Feedback  string     `json:"feedback" validate:"required"`
	Like      *bool      `json:"like"`
	CreatedBy string     `json:"created_by"`
	Page      string     `json:"page" validate:"required"`
	TenantId  string     `json:"tenant_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
