package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/yurianxdev/go-examples/gorm/models"
)

var db *gorm.DB

func main() {
	log.Printf("Intializing database...\n")
	var err error
	// Conecta a la base de datos y guarda su puntero en <db>, una variable a nivel de
	// paquete.
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=password sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})

	log.Printf("Database runing\n")
	log.Printf("Initializing the server...\n")

	// Handler para users/create, adentro se encuentra la logica para determinar su metodo,
	// pero una solucion mas elegante podria hacerse usando un multiplexer de terceros o algun
	// framework.
	http.Handle("/users/create", http.HandlerFunc(users))

	// Escucha conexiones y las acepta con el handler por default (por eso el nil).
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error serving: %v\n", err)
	}
}

func users(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request Income: %v\n", req)
	// Retorna un bad request si el method no es POST.
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		log.Printf("Responsed Error\n")
		return
	}

	var user models.User

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	// Decodifica el body del request en un User.
	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Printf("Responsed Error: %v\n", err)
		return
	}

	// Crea el usuario en la base de datos.
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Printf("Responsed Error: %v\n", err)
		return
	}

	// Crea JSON con la respuesta exitosa
	res, err := json.Marshal(models.ResponseSucceed{Message: "user created"})

	log.Printf("Created user: %v", user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}
