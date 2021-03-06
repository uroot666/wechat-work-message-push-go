package controller

import (
	"github.com/cloverzrg/wechat-work-message-push-go/grafana"
	"github.com/cloverzrg/wechat-work-message-push-go/qyapi"
	"github.com/gin-gonic/gin"
)

// GrafaneHandler 回调函数
func GrafaneHandler(c *gin.Context) {
	noti := grafana.Notification{}
	err := c.BindJSON(&noti)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	err = qyapi.SendTextGrafanaMessage(noti)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "ok")
}
