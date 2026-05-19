package helper

import "github.com/gofiber/fiber/v2"

// ResponseFormat defines the standard API response structure
type ResponseFormat struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse sends a success response
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(ResponseFormat{
		Status: "success",
		Data:   data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ResponseFormat{
		Status:  "error",
		Message: message,
	})
}
