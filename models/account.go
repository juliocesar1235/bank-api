package models

import (
	"time"
	"net/http"
	"fmt"
	"github.com/google/uuid"
)

type Account struct {
	ID 					uuid.UUID `json:"acc_id"`
	Type 				string 		`json:"acc_type"`
	Level				string 		`json:"acc_level"`
	Clabe 			string 		`json:"acc_clabe"`
	CustomerId 	string  	`json:"acc_customer_id"`
	CreatedAt 	time.Time `json:"acc_created_at"`
	UpdatedAt 	time.Time `json:"acc_updated_at"`
	DeletedAt 	time.Time `json:"acc_deleted_at"`
}

func (i *Account) Bind(r *http.Request) error {
	if i.Type == "" {
			return fmt.Errorf("Type is a required field")
	}
	if i.Level == "" {
			return fmt.Errorf("Level is a required field")
	}
	if i.CustomerId == "" {
		return fmt.Errorf("CustomerId is a required field")
	}

	return nil
}

func (*Account) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}