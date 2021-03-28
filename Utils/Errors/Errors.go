package Errors

type ResError struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
