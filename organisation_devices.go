package main

import "github.com/gin-gonic/gin"

func attachDevice(c *gin.Context) {
	c.String(200, "Attach Device Function")
}
