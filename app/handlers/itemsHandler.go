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

func CreateProduct(ctx *gin.Context) {
	var (
		item   models.Items
		err    error
		dbConn *sql.DB
	)

	if err = ctx.ShouldBindJSON(&item); err != nil {
		fmt.Println(item)
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}

	defer dbConn.Close()
	if err = db.InsertItem(dbConn, item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Item agregado": item})
}

func DeleteItem(ctx *gin.Context) {
	var (
		dbConn *sql.DB
		err    error
	)
	idDel, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 0": err})
		return
	}
	dbConn, err = db.GetDBConnection()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 1": err})
		return
	}
	defer dbConn.Close()

	if err = db.DeleteItem(dbConn, idDel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}
	mensaje := "Item " + fmt.Sprint(idDel) + " eliminado"
	ctx.JSON(200, gin.H{"message": mensaje})
}

func UpdatePrecio(ctx *gin.Context) {
	var (
		dbConn *sql.DB
		err    error
		item   models.Items
	)
	idUpdate, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err})
		return
	}
	if err = ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}
	defer dbConn.Close()
	precio, _ := strconv.ParseFloat(item.Precio, 64)
	if err = db.ActPrecio(dbConn, idUpdate, precio); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"errror": err})
		return
	}
	mensaje := "Precio del item " + fmt.Sprint(idUpdate) + " actualizado"
	ctx.JSON(http.StatusOK, gin.H{"message ": mensaje})

}

func UpdateStock(ctx *gin.Context) {
	var (
		dbConn *sql.DB
		err    error
		item   models.Items
	)
	idUpdate, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err})
		return
	}
	if err = ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 2": err})
		return
	}

	dbConn, err = db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 3": err})
		return
	}
	defer dbConn.Close()
	stock := item.Stock
	if err = db.ActStock(dbConn, idUpdate, stock); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	mensaje := "Stock del item " + fmt.Sprint(idUpdate) + " actualizado"
	ctx.JSON(http.StatusOK, gin.H{"message ": mensaje})

}

func GetItems(ctx *gin.Context) {
	var items [][]string
	dbConn, err := db.GetDBConnection()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error ": err})
		return
	}
	defer dbConn.Close()
	if items, err = db.GetItems(dbConn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 1": err})
		return
	}

	ctx.JSON(201, gin.H{"Items": items})
}
