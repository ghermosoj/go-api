package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ghermosoj/go-api/internal/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        api.GET("/items", handlers.GetItems)
        api.GET("/items/:id", handlers.GetItem)
        api.POST("/items", handlers.CreateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)
    }

    return r
}
