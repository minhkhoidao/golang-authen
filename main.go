package main

import (
	"golang-authen/config"
	"golang-authen/controllers"
	"golang-authen/driver"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	driver.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run() // listen and serve on 0.0.0.0:8080
}
