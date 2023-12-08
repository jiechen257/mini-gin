package mock

import (
	"mini-gin/gin"
	"net/http"
)

func MockRecovery() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello ginktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gin.Context) {
		names := []string{"ginktutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
