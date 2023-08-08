package enumerations

type Action string

var (
	ActionSubmitted Action = "submitted"
	ActionCompleted Action = "completed"
	ActionFailed    Action = "failed"
)

func (t Action) String() string {
	return string(t)
}
