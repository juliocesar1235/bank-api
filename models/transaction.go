package models

import (
	"time"
	"net/http"
	"fmt"
	"github.com/google/uuid"
)

type Transaction struct {
	ID 							uint64 			`json:"trn_id"`
	DatePerformed 	time.Time 	`json:"trn_date_perfomed"`
	Type 						string 			`json:"trn_type"`
	Amount 					float64 		`json:"trn_amount"`
	AccountId 			uuid.UUID 	`json:"trn_account_id"`
	CreatedAt 			time.Time 	`json:"ctr_created_at"`
	UpdatedAt 			time.Time 	`json:"ctr_updated_at"`
	DeletedAt 			*time.Time 	`json:"ctr_deleted_at"`
}

type TransactionList struct {
	Transactions []Transaction `json:"transactions"`
}

func (i *Transaction) Bind(r *http.Request) error {
	if i.DatePerformed.IsZero() {
			return fmt.Errorf("name is a required field")
	}
	if i.Type == "" {
			return fmt.Errorf("Type is a required field")
	}
	if i.Amount == 0.0 {
		return fmt.Errorf("Amount is a required field")
	}
	if i.AccountId.String() == "" {
		return fmt.Errorf("AccountId is a required field")
	}
	return nil
}

func (*TransactionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Transaction) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}