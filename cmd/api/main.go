package main

import (
	"log"

	"github.com/mateusfaustino/go-rest-api-i/configurations"
	"github.com/mateusfaustino/go-rest-api-i/routes"
)

func main() {
	connection, err := configurations.ConnectDb()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer connection.Close()

	

	server := routes.SetupRouter(connection)
	log.Println("Servidor rodando na porta 8080")
	server.Run(":8080")
}
