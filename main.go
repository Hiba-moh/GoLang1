package main

import (
	"test/golang/config"

	"github.com/gin-gonic/gin"
)

type User struct {
	user string
}

func main() {
	config.Conf()
	var users User
	users.user = "Your name"
	server := gin.Default()
	server.LoadHTMLGlob("./Templates/*")
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.tmpl", map[string]string{"name": users.user, "t1": config.Title})
	})
	server.POST("/testapi", func(c *gin.Context) {
		users.user = c.DefaultQuery("name", "unknown")
		c.String(200, "Hello %s", users.user)
	})
	server.Run(":" + config.Port)
}
