package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/jagadeeshtakkalaki/gin/controllers"
	"log"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
		log.Printf("Registering /ping route")
            api.GET("/ping", controllers.Ping)

            examples := api.Group("/examples")
            {
                examples.GET("", controllers.GetExamples)
                examples.POST("", controllers.CreateExample)
                examples.GET(":id", controllers.GetExample)
                examples.PUT(":id", controllers.UpdateExample)
                examples.DELETE(":id", controllers.DeleteExample)
            }
    }
}
