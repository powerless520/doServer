package route

import "github.com/gin-gonic/gin"

func Init(app *gin.Engine) {

	rootGroup := app.Group("/")
	{
		rootGroup.GET("/status", func(context *gin.Context) {
			context.JSON(200, "status")
		})
	}

	apiGroup := app.Group("/api")
	{
		apiGroup.GET("/login", func(context *gin.Context) {
			context.String(200, "login")
		})
	}
}
