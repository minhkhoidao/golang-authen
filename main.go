package main

import (
	"golang-authen/config"
	"golang-authen/controllers"
	"golang-authen/driver"
	"golang-authen/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	driver.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
