package db

import (
	"app/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Insert(db *sql.DB, usuario models.User) error {
	query := "INSERT INTO user (nombre, email) VALUE (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(usuario.Nombre, usuario.Email); err != nil {
		return fmt.Errorf("error al ejecutar la consulta")
	}
	fmt.Println("Usuario agregado correctamente")
	return nil
}

func UpdateName(db *sql.DB, idMod int, nombreNuevo string) error {
	if nombreNuevo == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	query := "UPDATE user set nombre = ? where idUser = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	fmt.Println(nombreNuevo)
	defer stmt.Close()
	if _, err := stmt.Exec(nombreNuevo, idMod); err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	fmt.Println("Usuario modificado")
	return nil
}

func UpdateEmail(db *sql.DB, idMod int, emailMod string) error {
	if emailMod == "" {
		return errors.New("el email no puede estar vacio")
	}
	query := "update user set email = ? where idUser = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	if _, err := stmt.Exec(emailMod, idMod); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println("Email modificado")
	return nil
}

func DeleteUser(db *sql.DB, idDelete int) error {
	query := "DELETE from user where idUser = ?"
	stmt, err := db.Prepare(query)

	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	fmt.Println(idDelete)
	if _, err := stmt.Exec(idDelete); err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Printf("Usuario %v eliminado", idDelete)
	return nil
}

func GetUser(db *sql.DB, id int) (string, error) {
	query := "SELECT nombre FROM user where idUser = ? "
	stmt, err := db.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	var user string
	if err := stmt.QueryRow(id).Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("%s", err)
		}
		return "", fmt.Errorf("%s", err)
	}
	return user, nil
}

func GetUsers(db *sql.DB) ([][]string, error) {
	query := "SELECT * FROM user"
	stmt, err := db.Prepare(query)
	if err != nil {
		return [][]string{}, fmt.Errorf("%s", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return [][]string{}, fmt.Errorf("%s", err)
	}
	defer rows.Close()
	var users [][]string
	for rows.Next() {
		var id, user, email string
		if err := rows.Scan(&id, &user, &email); err != nil {
			return [][]string{}, fmt.Errorf("%s", err)
		}

		users = append(users, []string{id, user, email})
	}
	return users, nil
}
