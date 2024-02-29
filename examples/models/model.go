package models

import (
	"encoding/json"
	"time"
)

// User is a model for the users table
type User struct {
	ID          string          `docstore:"id"`
	Name        string          `docstore:"name"`
	Age         uint            `docstore:"age"`
	DateJoiend  time.Time       `docstore:"date_joined"`
	IsAdmin     bool            `docstore:"is_admin"`
	Preferences json.RawMessage `docstore:"preferences"`
	Expenditure float64         `docstore:"expenditure"`
}

// TableName returns the table name for the model
func (User) TableName() string { return "users" }
