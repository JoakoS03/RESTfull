package db

import (
	"app/models"
	"database/sql"
	"fmt"
)

func CreateCompra(db *sql.DB, compra models.Compra) (string, error) {
	query := "INSERT INTO compra (idItem, idUser, precioTotal, cant) VALUE (?,?,?,?)"
	stmt, err := db.Prepare(query)

	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(compra.IdItem, compra.IdUser, compra.PrecioTotal, compra.Cant)
	if err != nil {
		return "", fmt.Errorf("error ejecutando la consulta: %w", err)
	}

	lastInserted, err := res.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	compraStr := fmt.Sprint(lastInserted)

	fmt.Println("Compra agregada")
	return compraStr, nil
}

func GetCompraUser(db *sql.DB, idUser int) ([][]string, error) {
	query := "SELECT item.idItem, compra.idUser, compra.precioTotal, compra.cant FROM compra inner join item on compra.idItem = item.idItem INNER join user on user.idUser = compra.idUser where compra.idUser = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return [][]string{}, fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	var compras [][]string

	rows, err := stmt.Query(idUser)
	if err != nil {
		fmt.Println(err)
		return [][]string{}, fmt.Errorf("%s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var idItem, idUser, precio, cant string
		if err := rows.Scan(&idItem, &idUser, &precio, &cant); err != nil {
			fmt.Println(err)
			return [][]string{}, fmt.Errorf("%s", err)
		}
		compras = append(compras, []string{idItem, idUser, precio, cant})
	}
	return compras, nil
}
