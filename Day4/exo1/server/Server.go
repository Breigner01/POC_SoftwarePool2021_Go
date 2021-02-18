package server

import (
	"SoftwareGoDay2/routes"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	routes.ApplyRoutes(r)
	return r
}
