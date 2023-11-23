package cmd

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func (controller Controller) AutoReloadData(c *gin.Context) {
	err := controller.service.AutoReload()
	if err != nil {
		if err != nil {
			generateResponse(c, 5900, gin.H{
				"message": "internal error",
			})
		}
	}
	return
}

func generateResponse(c *gin.Context, code int, resp interface{}) {
	c.JSON(code, resp)
}
