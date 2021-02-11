package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
}
