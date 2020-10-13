package api

import (
	"fmt"

	"github.com/cloverzrg/wechat-work-message-push-go/logger"
	"github.com/gin-gonic/gin"
)

// Start 启动webhook
func Start() (err error) {
	r := gin.Default()
	r.Use(logger.GinLogger())
	SetRoute(r)
	fmt.Println("启动webhook...")
	err = r.Run("0.0.0.0:5555")
	return err
}

func init() {
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
}
