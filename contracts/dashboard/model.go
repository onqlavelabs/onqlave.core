package dashboard

import (
	"time"
)

type Operations struct {
	Insight Insight `json:"insight"`
	Events  []Event `json:"events"`
}

type Insight struct {
	NumberOfArx          int `json:"number_of_arx"`
	NumberOfApplications int `json:"number_of_applications"`
	NumberOfApiKeys      int `json:"number_of_apikeys"`
	NumberOfUsers        int `json:"number_of_users"`
}

type Event struct {
	Id         string    `json:"id"`
	Operation  string    `json:"operation"`
	Message    string    `json:"message"`
	IsError    bool      `json:"is_error"`
	CreatedAt  time.Time `json:"created_at"`
	ObjectType string    `json:"object_type"`
	ObjectName string    `json:"object_name"`
	UserName   string    `json:"user_name"`
	UserID     string    `json:"user_id"`
}
