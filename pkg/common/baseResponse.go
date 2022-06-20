package common

type BaseErrorResponse struct {
	Error struct {
		ErrorCode    int32  `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"error"`
}

func NewBaseErrorResponse(errorCode int32, errorMessage string) BaseErrorResponse {
	e := &BaseErrorResponse{Error: struct {
		ErrorCode    int32  "json:\"error_code\""
		ErrorMessage string "json:\"error_message\""
	}{ErrorCode: errorCode, ErrorMessage: errorMessage}}
	return *e
}
