package controller

import "github.com/gin-gonic/gin"

// Index 首页调整
func Index(c *gin.Context) {
	c.String(400, "post to /push")
}
