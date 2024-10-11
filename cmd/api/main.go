package main

import (
	"log"

	"github.com/mateusfaustino/go-rest-api-i/config"
	"github.com/mateusfaustino/go-rest-api-i/routes"
)

func main() {
	connection, err := config.ConnectDb()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer connection.Close()

	

	router := routes.SetupRouter(connection)
	log.Println("Servidor rodando na porta 8080")
	router.Run(":8080")
}
