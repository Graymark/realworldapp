package server

import (
	"github.com/Graymark/realworldapp/uc"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	interactor uc.Interactor
}

func NewRouterHandler(i uc.Interactor) RouterHandler {
	return RouterHandler{
		interactor: i,
	}
}

func (rH RouterHandler) SetRoutes(router *gin.Engine) {
	router.GET("/health", rH.health)
}
