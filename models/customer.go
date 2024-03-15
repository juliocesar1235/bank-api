package models

import (
	"time"
	"net/http"
	"fmt"
	"github.com/google/uuid"
)

type Customer struct {
	ID 				uuid.UUID `json:"ctr_id"`
	Name 			string 		`json:"ctr_name"`
	FirstName string 		`json:"ctr_first_name"`
	LastName 	string 		`json:"ctr_last_name"`
	Email 		string  	`json:"ctr_email"`
	Phone 		string		`json:"ctr_phone"`
	Dob				time.Time	`json:"ctr_dob"`
	CreatedAt time.Time `json:"ctr_created_at"`
	UpdatedAt time.Time `json:"ctr_updated_at"`
	DeletedAt time.Time `json:"ctr_deleted_at"`
}

func (i *Customer) Bind(r *http.Request) error {
	if i.Name == "" {
			return fmt.Errorf("name is a required field")
	}
	if i.FirstName == "" {
			return fmt.Errorf("FirstName is a required field")
	}
	if i.LastName == "" {
		return fmt.Errorf("LastName is a required field")
	}
	if i.Email == "" {
		return fmt.Errorf("Email is a required field")
	}
	if i.Phone == "" {
		return fmt.Errorf("Phone is a required field")
	}
	return nil
}

func (*Customer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}