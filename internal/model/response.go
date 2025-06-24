package model

type SuccessRespones struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorRespones struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Errors     error  `json:"error" `
}
