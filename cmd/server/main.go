package main

import "github.com/ghermosoj/go-api/internal/routes"


func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}