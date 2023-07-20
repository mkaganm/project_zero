package services

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type ErrorResponse struct {
	Status       string `json:"status"`
	Error        error  `json:"error"`
	ErrorMessage string `json:"errorMessage"`
}
