package routes

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/createUser", handlers.CreateUser)

		//Actulizar Nombre
		userGroup.PUT("/actUserName/:id/:nombre", handlers.UpdateName)

		//Actualizar email
		userGroup.PUT("/actEmail/:id", handlers.UpdateEmail)

		//Borrar un usuario
		userGroup.DELETE("/delUser/:idDel", handlers.DeleteUser)

		//Devolver un usuario
		userGroup.GET("/getUser/:idU", handlers.GetUser)

		//Devolver todos los usuario
		userGroup.GET("/getUsers", handlers.GetUsers)
	}
}
