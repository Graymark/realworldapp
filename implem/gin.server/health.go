package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) health(c *gin.Context) {

	if err := rH.interactor.Health(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"hey": "I'm boring !"})
}

