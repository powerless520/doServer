package main

import (
	"doServer/route"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("demo")

	app := gin.Default()

	route.Init(app)
	_ = app.Run(":9090")
}
