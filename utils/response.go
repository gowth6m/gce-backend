package utils

import (
	// "encoding/json"
	"greatcomcatengineering.com/backend/models"
	// "net/http"
	"github.com/gin-gonic/gin"
)

// // RespondWithJSON sends a JSON response with a generic content type.
// func RespondWithJSON[T any](w http.ResponseWriter, status int, message string, content T) {
// 	response := models.ApiResponse[T]{
// 		Status:  status,
// 		Message: message,
// 		Data: content,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(response)
// }

// // RespondWithError is a convenience wrapper for sending error messages using the ApiResponse structure.
//
//	func RespondWithError(w http.ResponseWriter, status int, message string) {
//		RespondWithJSON[interface{}](w, status, message, nil)
//	}

func RespondWithJSON[T any](c *gin.Context, status int, message string, content T) {
	response := models.ApiResponse[T]{
		Status:  status,
		Message: message,
		Data:    content,
	}

	c.JSON(status, response)
}

// RespondWithError is a convenience wrapper for sending error messages using the ApiResponse structure.
func RespondWithError(c *gin.Context, status int, message string) {
	RespondWithJSON[interface{}](c, status, message, nil)
	c.Abort() // Ensures no further handlers are called after this response
}
