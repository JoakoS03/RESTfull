package routes

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCompraRoutes(r *gin.Engine) {
	userGroup := r.Group("/compra")
	{
		userGroup.POST("/crearCompra", handlers.CreateCompra)
		userGroup.GET("/compraUser/:idUser", handlers.GetCompraUser)
	}
}
