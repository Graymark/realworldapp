package main

import (
	"github.com/Graymark/realworldapp/implem/gin.server"
	"github.com/Graymark/realworldapp/implem/logrus.logger"
	"github.com/Graymark/realworldapp/infra"
	"github.com/Graymark/realworldapp/uc"
)

func main() {
	ginServer := infra.NewServer(8082, infra.DebugMode)

	rH := server.NewRouterHandler(uc.NewInteractor(logger.NewLogger("TEST", "podName")))

	rH.SetRoutes(ginServer.Router)

	ginServer.Start()
}
