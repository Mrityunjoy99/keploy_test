package entity

import "github.com/sample_project/src/infrastructure/database"

type User struct {
	database.BaseModel
	FirstName string
	LastName  string
	Email     string
}
