package env

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Route(routes gin.IRoutes) {
	routes.GET("/env", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, os.Environ())
	})
}
