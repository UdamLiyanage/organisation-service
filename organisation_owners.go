package main

import "github.com/gin-gonic/gin"

func attachUser(c *gin.Context) {
	c.JSON(200, "Attach Organisation Function")
}
