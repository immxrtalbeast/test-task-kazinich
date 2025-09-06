package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immxrtalbeast/rtp-multiplier/internal/service"
)

type RTPController struct {
	service service.RTPMultiplierService
}

func NewRTPController(service service.RTPMultiplierService) *RTPController {
	return &RTPController{service: service}
}

func (c *RTPController) GetMultiplier(ctx *gin.Context) {
	multiplier := c.service.GenerateMultiplier()
	ctx.JSON(http.StatusOK, gin.H{
		"result": multiplier,
	})
}
