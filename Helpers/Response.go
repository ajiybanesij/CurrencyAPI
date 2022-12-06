package Helpers

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w *gin.Context, status bool, message string, payload interface{}, statusCode int) {
	w.Writer.Header().Set("Content-Type", "application/json")

	res := ResponseData{
		Status:  status,
		Message: message,
		Data:    payload,
	}

	w.JSON(statusCode, res)

	return
}
