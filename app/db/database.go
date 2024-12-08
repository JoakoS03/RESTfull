package db

import (
	"database/sql"
	"fmt"
)

func GetDBConnection() (*sql.DB, error) {
	dns := "root:@tcp(localhost:3307)/prueba"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexion: %w", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("erro al verificar la conexion: %w", err)
	}
	return db, nil
}
