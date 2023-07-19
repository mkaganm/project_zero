package services

type ErrorResponse struct {
	Status       string `json:"status"`
	Error        error  `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}
