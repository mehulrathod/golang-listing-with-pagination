package Routers

import (
	"github.com/gin-gonic/gin"
	Controller "testTask/controllers/api/v1"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	r.Static("/static", "static")
	r.Static("/templates", "templates")
	r.LoadHTMLGlob("templates/*")


	v1 := r.Group("/api/v1")

	v1.POST("employee-list", Controller.EmployeesList)

	return r

}