package models

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email" gorm:"type:varchar(100);unique_index"`
	Role  string `json:"role" gorm:"size:255"`
}
