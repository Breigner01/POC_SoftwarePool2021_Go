package routes

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func world(c *gin.Context) {
	str := os.Getenv("HELLO_MESSAGE")

	if str == "" {
		c.String(http.StatusNotFound, "No message defined")
	} else {
		c.String(http.StatusOK, str)
	}
}

func repeatMyQuery(c *gin.Context) {
	str, status := c.GetQuery("message")

	if !status || str == "" {
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
		c.String(http.StatusBadRequest, "An error has occurred: %w", err)
	} else if str == "" {
		c.String(http.StatusBadRequest, "Missing message")
	} else {
		c.String(http.StatusOK, str)
	}
}

func repeatAllMyQueries(c *gin.Context) {

}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/health", health)
	r.GET("/hello", world)
	r.GET("/repeat-my-query", repeatMyQuery)
	r.GET("/repeat-my-param/:message", repeatMyParams)
	r.POST("/repeat-my-body", repeatMyBody)
	r.GET("/repeat-my-header", repeatMyHeader)
	r.GET("/repeat-my-cookie", repeatMyCookie)
	r.GET("/repeat-all-my-queries", repeatAllMyQueries)
}
