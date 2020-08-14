package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/yurianxdev/go-examples/gorm/controllers"
	"github.com/yurianxdev/go-examples/gorm/db"
	"github.com/yurianxdev/go-examples/gorm/models"
)

func main() {
	log.Printf("Intializing database...\n")
	var err error
	// Conecta a la base de datos y guarda su puntero en <db>, una variable a nivel de
	// paquete.
	db.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=password sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.DB.Close()
	db.DB.AutoMigrate(&models.User{})

	log.Printf("Database runing\n")
	log.Printf("Initializing the server...\n")

	// Handler para users/create, adentro se encuentra la logica para determinar su metodo,
	// pero una solucion mas elegante podria hacerse usando un multiplexer de terceros o algun
	// framework.
	http.Handle("/users/create", http.HandlerFunc(controllers.Users))

	// Escucha conexiones y las acepta con el handler por default (por eso el nil).
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error serving: %v\n", err)
	}
}
