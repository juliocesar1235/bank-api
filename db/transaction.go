package db
import (
    "time"
    "github.com/juliocesar1235/bank-api/models"
)
func (db Database) GetTransactionsByAccountId(accId string) (*models.TransactionList, error) {
    list := &models.TransactionList{}
		query := `SELECT * FROM transaction_acc WHERE trn_account_id = $1 ORDER BY ID DESC`
    rows, err := db.Conn.Query(query, accId)
    if err != nil {
        return list, err
    }
    for rows.Next() {
        var transaction models.Transaction
        err := rows.Scan(
					&transaction.ID, 
					&transaction.DatePerformed, 
					&transaction.Type, 
					&transaction.Amount, 
					&transaction.AccountId,
					&transaction.CreatedAt,
					&transaction.UpdatedAt,
					&transaction.DeletedAt)
        if err != nil {
            return list, err
        }
        list.Transactions = append(list.Transactions, transaction)
    }
    return list, nil
}
func (db Database) createTransaction(transaction *models.Transaction) error {
    var id uint64
    var createdAt time.Time
    query := `INSERT INTO transaction_acc (trn_date_performed,trn_type,trn_amount,trn_account_id) VALUES ($1,$2,$3,$4) RETURNING trn_id, trn_created_at`
    err := db.Conn.QueryRow(query, transaction.DatePerformed, transaction.Type, transaction.Amount, transaction.AccountId).Scan(&id, &createdAt)
    if err != nil {
        return err
    }
    transaction.ID = id
    transaction.CreatedAt = createdAt
    return nil
}
