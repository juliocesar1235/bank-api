
CREATE TABLE IF NOT EXISTS transaction_acc(
  trn_id BIGSERIAL,
  trn_date_perfomed DATE NOT NULL,
  trn_type VARCHAR(15),
  trn_amount numeric NOT NULL,
  trn_account_id uuid NOT NULL,
  trn_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  trn_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  trn_deleted_at TIMESTAMP,
  PRIMARY KEY(trn_id),
  CONSTRAINT fk_account FOREIGN KEY(trn_account_id) REFERENCES account(acc_id)
);