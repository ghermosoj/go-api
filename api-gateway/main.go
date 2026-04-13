package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. DEFINICIÓN DE SEGURIDAD
	const MiClaveSecreta = "token-super-seguro-2024"
	const DestinoInterno = "http://localhost:3000"

	// 2. MIDDLEWARE DE VALIDACIÓN
	r.Use(func(c *gin.Context) {
		// Buscamos un header llamado "Authorization"
		key := c.GetHeader("X-Gateway-Auth")

		if key != MiClaveSecreta {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensaje": "No tienes permiso para pasar por el Gateway",
			})
			c.Abort()
			return
		}
		c.Next()
	})

	// 3. PROXY INVERSO (El puente)
	r.Any("/*proxyPath", func(c *gin.Context) {
		remote, _ := url.Parse(DestinoInterno)
		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})

	fmt.Println("🛡️ Gateway protegiendo en el puerto :8080")
	r.Run(":8080")
}