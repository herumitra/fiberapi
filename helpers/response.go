package helpers

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse is a struct to format error message
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
