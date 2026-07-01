package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:  "ok",
		Message: "Diary Of Aditya API is running",
		Version: "1.0.0",
	})
}
