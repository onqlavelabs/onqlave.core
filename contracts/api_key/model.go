package api_key

import (
	"github.com/onqlavelabs/onqlave.core/enumerations"
	"time"

	"github.com/onqlavelabs/onqlave.core/contracts/acl"
	"github.com/onqlavelabs/onqlave.core/contracts/application"
	arx "github.com/onqlavelabs/onqlave.core/contracts/arx"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
)

type APIKey struct {
	ID            string                    `json:"id"`
	AccessKey     string                    `json:"access_key"`
	CreatedAt     string                    `json:"created_at"`
	Status        string                    `json:"status"`
	CreatedBy     *common.ShortUserInfo     `json:"created_by"`
	ApplicationID string                    `json:"application_id,omitempty"`
	Application   *common.ShortResourceInfo `json:"application"`
	ArxID         string                    `json:"cluster_id,omitempty"`
	Arx           *common.ShortResourceInfo `json:"cluster"`
	ACL           acl.ACL                   `json:"acl"`
	ArxUrl        string                    `json:"arx_url"`
}

type APIKeys struct {
	ACL      acl.ACL  `json:"acl"`
	APIKeys  []APIKey `json:"api_keys"`
	Model    Models   `json:"model"`
	Insights Insights `json:"insights"`
}

type SensitiveData struct {
	ID                    string     `json:"id"`
	AccessKey             string     `json:"access_key"`
	Status                string     `json:"status"`
	ClientKey             string     `json:"client_key"`
	ServerSigningKey      string     `json:"server_signing_key"`
	ServerCryptoAccessKey string     `json:"server_crypto_access_key"`
	ArxUrl                string     `json:"arx_url"`
	ProvidedAt            *time.Time `json:"provided_at"`
}

type Models struct {
	Applications []Application `json:"applications"`
	Arx          []Arx         `json:"clusters"`
}

type Application struct {
	common.ShortResourceInfo `json:",inline"`
	Label                    string                 `json:"label"`
	ApplicationTechnology    application.Technology `json:"application_technology"`
}

type Arx struct {
	common.ShortResourceInfo `json:",inline"`
	Purpose                  arx.Purpose                 `json:"purpose"`
	Plan                     arx.Plan                    `json:"plan"`
	Provider                 arx.Provider                `json:"provider"`
	Regions                  []arx.Region                `json:"regions"`
	Encryption               arx.EncryptionMethod        `json:"encryption"`
	RotationCycle            arx.EncryptionRotationCycle `json:"rotation_cycle"`
	CreatedBy                *common.ShortUserInfo       `json:"created_by"`
}

type Insights struct {
	TotalKeys    int64 `json:"total_keys"`
	TotalActive  int64 `json:"total_active"`
	TotalDeleted int64 `json:"total_deleted"`
}

type CreateAPIKey struct {
	ApplicationID         string `json:"application_id" validate:"required"`
	ClusterID             string `json:"cluster_id" validate:"required"`
	ApplicationTechnology string `json:"application_technology" validate:"required"`
}

type APIKeyVersion struct {
	Number    string    `json:"version_number" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
}

type EventAPIKeyStatusChange struct {
	KeyID      string                    `json:"key_id"`
	AccessKey  string                    `json:"access_key"`
	Status     enumerations.ApiKeyStatus `json:"status"`
	ProvidedAt *time.Time                `json:"provided_at"`
	Message    string                    `json:"message"`
	ArxUrl     string                    `json:"arx_url"`
	IsError    bool                      `json:"is_error"`
}

func (keys *APIKeys) SetACL(acl acl.ACL) {
	keys.ACL = acl
}

func (keys *APIKeys) SetAPIKey(apiKey APIKey) {
	keys.APIKeys = append(keys.APIKeys, apiKey)
}

func (keys *APIKeys) SetModel(apps []Application, clusters []Arx) {
	if apps != nil {
		keys.Model.Applications = make([]Application, len(apps))
		copy(keys.Model.Applications, apps)
	}

	if clusters != nil {
		keys.Model.Arx = make([]Arx, len(clusters))
		copy(keys.Model.Arx, clusters)
	}
}

func (keys *APIKeys) SetInsights(insights Insights) {
	keys.Insights = insights
}

func (app *Application) SetTechnology(applicationTechnology application.Technology) {
	app.ApplicationTechnology = applicationTechnology
}

func (c *Arx) SetPurpose(purpose arx.Purpose) {
	c.Purpose = purpose
}

func (c *Arx) SetPlan(plan arx.Plan) {
	c.Plan = plan
}

func (c *Arx) SetProvider(provider arx.Provider) {
	c.Provider = provider
}

func (c *Arx) SetRegions(regions []arx.Region) {
	if regions != nil {
		c.Regions = make([]arx.Region, len(regions))
		copy(c.Regions, regions)
	}
}

func (c *Arx) SetEncryption(encryption arx.EncryptionMethod) {
	c.Encryption = encryption
}

func (c *Arx) SetRotationCycle(rotationCycle arx.EncryptionRotationCycle) {
	c.RotationCycle = rotationCycle
}

func (c *Arx) SetCreatedBy(owner *common.ShortUserInfo) {
	c.CreatedBy = owner
}
