package routes

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine) {
	userGroup := r.Group("/item")
	{
		userGroup.POST("/addItem", handlers.CreateProduct)
		userGroup.DELETE("/delItem/:id", handlers.DeleteItem)
		userGroup.PUT("/actPrecio/:id", handlers.UpdatePrecio)
		userGroup.PUT("/actStock/:id", handlers.UpdateStock)
		userGroup.GET("/getItems", handlers.GetItems)
	}
}
