package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "1234"
	dbname   = "go_rest_api"
)

func ConnectDb() (*sql.DB, error) {
	// Define a string de conexão, incluindo a porta
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão está funcionando corretamente
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db, nil
}
