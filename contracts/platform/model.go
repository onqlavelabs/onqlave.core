package platform

type Status struct {
	Year          string `json:"year"`
	WebLink       string `json:"web_link"`
	SystemLink    string `json:"system_link"`
	SystemHealthy bool   `json:"system_healthy"`
	FeedbackLink  string `json:"feedback_link"`
}
