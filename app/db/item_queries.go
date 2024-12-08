package db

import (
	"app/models"
	"database/sql"
	"fmt"
)

func InsertItem(db *sql.DB, item models.Items) error {
	query := "INSERT INTO item (nombre, precio, stock) VALUE (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(item.Nombre, item.Precio, item.Stock); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println("item agregado")
	return nil
}

func DeleteItem(db *sql.DB, idDel int) error {
	query := "DELETE FROM item where idItem = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(idDel); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println("item borrado")
	return nil
}

func ActPrecio(db *sql.DB, idAct int, precioNuevo float64) error {
	query := "UPDATE item SET precio = ? where idItem = ? "
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(idAct, precioNuevo)
	defer stmt.Close()
	var itemNuevo sql.Result
	if itemNuevo, err = stmt.Exec(precioNuevo, idAct); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(itemNuevo.LastInsertId())
	return nil
}

func ActStock(db *sql.DB, idAct int, stock uint) error {
	query := "UPDATE item SET stock = ? where idItem = ? "
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	fmt.Println(stock, idAct)
	var itemNuevo sql.Result
	if itemNuevo, err = stmt.Exec(stock, idAct); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(itemNuevo)
	return nil
}

func GetItems(db *sql.DB) ([][]string, error) {
	query := "SELECT * FROM item"
	stmt, err := db.Prepare(query)
	if err != nil {
		return [][]string{}, fmt.Errorf("%s", err)
	}
	defer stmt.Close()

	var items [][]string
	rows, err := stmt.Query()
	if err != nil {
		return [][]string{}, fmt.Errorf("%s", err)
	}

	for rows.Next() {
		var id, nombre, precio, stock string
		if err := rows.Scan(&id, &nombre, &precio, &stock); err != nil {
			return [][]string{}, fmt.Errorf("%s", err)
		}
		items = append(items, []string{id, nombre, precio, stock})
	}
	return items, nil
}
