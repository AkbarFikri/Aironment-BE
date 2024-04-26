package model

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

type ServiceResponse struct {
	Code    int
	Error   bool
	Message string
	Payload any
}
