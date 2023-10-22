package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LivenessProbe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

func ReadinessProbe(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}
