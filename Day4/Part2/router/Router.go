package router

import (
	"SoftwareGoDay4/softcrypto"
	"SoftwareGoDay4/user"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func signupSession(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	str := string(bytes)
	arr := strings.Split(str, "\n")
	var content [][]string
	var newUser user.User

	if err != nil {
		c.String(http.StatusInternalServerError, "An error has occurred: %v", err)
	} else if str == "" {
		c.String(http.StatusBadRequest, "Missing credentials")
	} else {
		for i := 0; i < len(arr); i++ {
			content = append(content, strings.Split(arr[i], "="))
			if content[i][0] == "email" {
				newUser.Email = content[i][1]
			} else if content[i][0] == "password" {
				newUser.Password = content[i][1]
			}
		}
		user.UserDB = append(user.UserDB, newUser)
		c.String(http.StatusOK, "User successfully created !")
	}
}

func signinSession(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	str := string(bytes)
	arr := strings.Split(str, "\n")
	var content [][]string
	var tmpUser user.User

	if err != nil {
		c.String(http.StatusInternalServerError, "An error has occurred: %v", err)
	} else if str == "" {
		c.String(http.StatusBadRequest, "Missing credentials")
	} else {
		for i := 0; i < len(arr); i++ {
			content = append(content, strings.Split(arr[i], "="))
			if content[i][0] == "email" {
				for j := 0; j < len(user.UserDB); j++ {
					if user.UserDB[j].Email == content[i][1] {
						tmpUser.Email = content[i][1]
					}
				}
			} else if content[i][0] == "password" {
				for j := 0; j < len(user.UserDB); j++ {
					if user.UserDB[j].Password == content[i][1] {
						tmpUser.Password = content[i][1]
					}
				}
			}
		}
		if tmpUser.Email != "" && tmpUser.Password != "" {
			c.SetCookie("token", softcrypto.Encrypt(tmpUser.Email, user.APISECRET), 24 * 60 * 60,
				"/signin-session", "localhost:8080", true, false)
			c.String(http.StatusOK, "User successfully logged in !")
		}
	}
}

func meSession(c *gin.Context) {
	str, err := c.Cookie("token")

	if err != nil {
		c.String(http.StatusInternalServerError, "An error has occurred: %v", err)
	} else if str == "" {
		c.String(http.StatusForbidden, "Missing cookie")
	} else {
		email := softcrypto.Decrypt(str, user.APISECRET)
		for i := 0; i < len(user.UserDB); i++ {
			if email == user.UserDB[i].Email {
				res := "email=" + user.UserDB[i].Email + "\npassword" + user.UserDB[i].Password
				c.String(http.StatusOK, res)
				return
			}
		}
		c.String(http.StatusUnauthorized, "You can't do that")
	}
}

func signupJwt(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	str := string(bytes)
	arr := strings.Split(str, "\n")
	var content [][]string
	var newUser user.User

	if err != nil {

	}
}

func ApplyRoutes(r *gin.Engine) {
	r.POST("/signup-session", signupSession)
	r.POST("/signin-session", signinSession)
	r.GET("/me-session", meSession)
	r.POST("/signup-jwt", signupJwt)
}