CREATE TABLE IF NOT EXISTS customer(
  ctr_id uuid DEFAULT gen_random_uuid(),
  ctr_name VARCHAR(255) NOT NULL,
  ctr_first_name VARCHAR(255) NOT NULL,
  crt_last_name VARCHAR(255),
  ctr_email VARCHAR(510) UNIQUE NOT NULL,
  ctr_phone VARCHAR(20),
  ctr_dob DATE NOT NULL,
  ctr_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ctr_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  ctr_deleted_at TIMESTAMP,
  PRIMARY KEY(ctr_id)
);

CREATE TABLE IF NOT EXISTS account(
  acc_id uuid DEFAULT gen_random_uuid(),
  acc_type VARCHAR(255) NOT NULL,
  acc_level VARCHAR(255) NOT NULL,
  acc_clabe VARCHAR(255),
  acc_customer_id uuid NOT NULL,
  acc_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  acc_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  acc_deleted_at TIMESTAMP,
  PRIMARY KEY(acc_id),
  CONSTRAINT fk_customer FOREIGN KEY(acc_customer_id) REFERENCES customer(ctr_id)
);