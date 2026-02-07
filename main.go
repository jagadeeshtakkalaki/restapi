package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/jagadeeshtakkalaki/gin/db"
    "github.com/jagadeeshtakkalaki/gin/middlewares"
    "github.com/jagadeeshtakkalaki/gin/models"
    "github.com/jagadeeshtakkalaki/gin/routes"
)

func main() {
    r := gin.Default()
    r.Use(middlewares.DummyLogger())

    if err := db.Init(""); err != nil {
        log.Fatal(err)
    }
    if err := db.DB.AutoMigrate(&models.Example{}); err != nil {
        log.Fatal(err)
    }

    routes.RegisterRoutes(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Starting server on port %s", port)

    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
