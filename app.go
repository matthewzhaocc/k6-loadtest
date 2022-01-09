// This is a simple webserver to load test against
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Hello string `json:"hello"`
}

func main() {
	GetServer().Run()
}

func GetServer() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", Ping)
	return router
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Hello: "world"})
}
