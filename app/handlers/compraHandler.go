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

func CreateCompra(ctx *gin.Context) {
	var (
		dbConn *sql.DB
		err    error
		compra models.Compra
	)

	if err = ctx.ShouldBindJSON(&compra); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}

	defer dbConn.Close()
	var c string
	if c, err = db.CreateCompra(dbConn, compra); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Compra: " + c})
}

func GetCompraUser(ctx *gin.Context) {
	var (
		dbConn *sql.DB
		err    error
	)
	idUser, err := strconv.Atoi(ctx.Params.ByName("idUser"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}
	defer dbConn.Close()
	var compras [][]string
	if compras, err = db.GetCompraUser(dbConn, idUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}

	header := "Compra de usuario " + fmt.Sprint(idUser)
	ctx.JSON(http.StatusOK, gin.H{header: compras})
}
