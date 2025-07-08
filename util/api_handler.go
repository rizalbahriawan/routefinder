package util

import "github.com/gin-gonic/gin"

func APIResponse(c *gin.Context, message string, code int, status string, data interface{}) {
	c.JSON(code,
		gin.H{
			"message": message,
			"code":    code,
			"status":  status,
			"data":    data,
		})
}

func APIListResponse(c *gin.Context, message string, code int, status string, data interface{}, totalData uint) {
	c.JSON(code,
		gin.H{
			"message": message,
			"code":    code,
			"status":  status,
			"data":    data,
			"total":   totalData,
		})
}
