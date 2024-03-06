package handler

import (
	"github.com/MamushevArup/jwt-auth/internal/service"
	"github.com/gin-gonic/gin"
	"time"
)

// Disclaimer all validation will be made in handler layer

var (
	maxAge = int(time.Now().Add(60 * 24 * time.Hour).Unix())
)

type RouterInitializer interface {
	InitRoute() *gin.Engine
}

func NewHandler(service *service.Service) RouterInitializer {
	return &handler{service: service}
}

func (h *handler) InitRoute() *gin.Engine {
	router := gin.Default()
	// grouping to easy navigation
	auth := router.Group("/auth")
	{
		// wait user guid
		auth.POST("/generate-token/:guid", h.createToken)
		// return new access token in body and refresh in cookies
		auth.POST("/refresh/:guid", h.refresh)
	}
	return router
}

type handler struct {
	service *service.Service
}
