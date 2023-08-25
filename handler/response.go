package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s", op),
		"data": data,
	})
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"errorCode"`
}