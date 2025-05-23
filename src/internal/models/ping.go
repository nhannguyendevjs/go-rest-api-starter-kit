package models

type PingResponse struct {
	Success bool   `json:"success,omitempty" example:"true"`
	Data    string `json:"data,omitempty" example:"pong"`
	Error   bool   `json:"error,omitempty" example:"true"`
	Message string `json:"message,omitempty" example:"Invalid request"`
}
