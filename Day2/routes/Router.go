package routes

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func world(c *gin.Context) {
	c.String(http.StatusOK, "world")
}

func repeatMyQuery(c *gin.Context) {
	str, status := c.GetQuery("message")

	if !status {
		c.String(http.StatusBadRequest, "Missing message")
	} else {
		c.String(http.StatusOK, str)
	}
}

func repeatMyParams(c *gin.Context) {
	str := c.Param("message")

	if str == "" {
		c.String(http.StatusBadRequest, "Missing message")
	} else {
		c.String(http.StatusOK, str)
	}
}

func repeatMyBody(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	str := string(bytes)

	if err != nil {
		c.String(http.StatusBadRequest, "An error has occurred: %w", err)
	} else if str == "" {
		c.String(http.StatusBadRequest, "Missing message")
	} else {
		c.String(http.StatusOK, str)
	}
}

func repeatMyHeader(c *gin.Context) {
	str := c.GetHeader("Message")

	if str == "" {
		c.String(http.StatusBadRequest, "Missing message")
	} else {
		c.String(http.StatusOK, str)
	}
}

func repeatMyCookie(c *gin.Context) {
	str, err := c.Cookie("message")

	if err != nil {
		c.String(http.StatusBadRequest, "Missing message")
	} else {
		c.String(http.StatusOK, str)
	}
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-my-query", repeatMyQuery)
	r.GET("/repeat-my-param/:message", repeatMyParams)
	r.POST("/repeat-my-body", repeatMyBody)
	r.GET("/repeat-my-header", repeatMyHeader)
	r.GET("/repeat-my-cookie", repeatMyCookie)
}
