package models

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email" gorm:"type:varchar(100);unique_index"` // Crea un unique_index cuando se cree el modelo
	Role  string `json:"role" gorm:"size:255"`
}
