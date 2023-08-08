package enumerations

type Object string

var (
	ObjectTypeArx         Object = "arx"
	ObjectTypeApiKey      Object = "api_key"
	ObjectTypeApplication Object = "application"
	ObjectTypeFeedback    Object = "feedback"
	ObjectTypeAccount     Object = "account"
	ObjectTypeUser        Object = "user"
)

func (t Object) String() string {
	return string(t)
}
