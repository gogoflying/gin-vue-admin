package servers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReportFormat(c *gin.Context, success bool, msg string, json gin.H) {
	// 开始时间
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"msg":     msg,
		"data":    json,
	})
}

func ReportFormatXML(c *gin.Context, msg string) {
	// 开始时间
	c.XML(http.StatusOK, msg)
}

func ReportFormatXMLEx(c *gin.Context, success bool, msg string, json gin.H) {
	// 开始时间
	c.XML(http.StatusOK, gin.H{
		"success": success,
		"msg":     msg,
		"data":    json,
	})
}
