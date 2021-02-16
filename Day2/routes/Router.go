package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func world(c *gin.Context) {
	c.String(http.StatusOK, "world")
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
}
