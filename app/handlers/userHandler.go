package handlers

import (
	"app/db"
	"app/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	dbConn *sql.DB
	err    error
)

func CreateUser(ctx *gin.Context) {
	var (
		usuario models.User
		err     error
	)
	err = ctx.ShouldBindJSON(&usuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	dbConn, err := db.GetDBConnection()
	if err != nil {
		fmt.Println("Error al conectar con la base de datos:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer dbConn.Close()
	if err = db.Insert(dbConn, usuario); err != nil {
		fmt.Print("hubo un error al insertar\n ")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": usuario.IdUser, "nombre": usuario.Nombre, "email": usuario.Email})
}

func UpdateName(ctx *gin.Context) {
	idMod, _ := strconv.Atoi(ctx.Params.ByName("id"))
	nameMod := ctx.Params.ByName("nombre")

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer dbConn.Close()
	if err = db.UpdateName(dbConn, idMod, nameMod); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado"})
}

func UpdateEmail(ctx *gin.Context) {
	idMod, err := strconv.Atoi(ctx.Params.ByName("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var usuario models.User //Para el bindeo no usar tipos primitivos

	err = ctx.ShouldBindJSON(&usuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err})
		return
	}

	dbConn, err := db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2 ": err})
		return
	}
	defer dbConn.Close()

	if err = db.UpdateEmail(dbConn, idMod, usuario.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}

	ctx.JSON(202, gin.H{"message": "Email actualizado"})
}

func DeleteUser(ctx *gin.Context) {
	idDel, err := strconv.Atoi(ctx.Params.ByName("idDel"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": " Dale hermano que hace" + err.Error()})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}
	defer dbConn.Close()
	err = db.DeleteUser(dbConn, idDel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "usuario borrado"})
}

func GetUser(ctx *gin.Context) {
	idUser, err := strconv.Atoi(ctx.Params.ByName("idU"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1 ": "Sos un inutil"})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}

	defer dbConn.Close()
	var usuario string
	if usuario, err = db.GetUser(dbConn, idUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 4": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"User ": usuario})
}

func GetUsers(ctx *gin.Context) {
	var usuarios [][]string
	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 1": err})
		return
	}
	defer dbConn.Close()
	if usuarios, err = db.GetUsers(dbConn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}

	ctx.JSON(201, gin.H{"Usuarios": usuarios})
}
