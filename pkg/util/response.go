package util

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func ErrorResponse(c *gin.Context, status int, err string) {
	c.JSON(status, gin.H{"error": err})
}
