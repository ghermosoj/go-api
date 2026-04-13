package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ghermosoj/go-api/internal/handlers"
    "net/http"
)

func GatewayGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
        
		const sharedSecret = "comunicacion-interna-segura-2026"
		
		if c.GetHeader("X-Internal-Secret") != sharedSecret {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Acceso denegado. Las peticiones deben pasar por el API Gateway.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

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
