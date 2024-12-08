package main

import (
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Rutas de usuario
	routes.RegisterUserRoutes(r)

	//Rutas de items
	routes.RegisterItemRoutes(r)

	//Rutas de compras
	routes.RegisterCompraRoutes(r)

	r.Run(":8080")
}
