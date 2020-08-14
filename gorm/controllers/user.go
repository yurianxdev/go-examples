package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yurianxdev/go-examples/gorm/db"
	"github.com/yurianxdev/go-examples/gorm/models"
)

func Users(w http.ResponseWriter, req *http.Request) {
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
	if err := db.DB.Create(&user).Error; err != nil {
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
