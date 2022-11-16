package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckModel struct {
	AppName string
	Version string
	Status  string
}

func HealthCheck(ctx *gin.Context) {
	appCheck := HealthCheckModel{
		AppName: "Sample toy app",
		Version: "0.0.1",
		Status:  "healthy",
	}
	ctx.JSON(http.StatusOK, appCheck)
}
