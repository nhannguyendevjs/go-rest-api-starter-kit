package models

type Book struct {
	ID     int    `json:"id" example:"1"`
	Title  string `json:"title" example:"The Go Programming Language"`
	Author string `json:"author" example:"Alan A. A. Donovan"`
}

type BookResponse struct {
	Success bool   `json:"success,omitempty" example:"true"`
	Data    []Book `json:"data,omitempty"`
	Error   bool   `json:"error,omitempty" example:"true"`
	Message string `json:"message,omitempty" example:"Invalid request"`
}
