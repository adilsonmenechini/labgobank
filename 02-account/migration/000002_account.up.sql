CREATE TABLE IF NOT EXISTS "accounts" (
  "account_number" VARCHAR(255) PRIMARY KEY,
  "account_type" VARCHAR(255) NOT NULL,
  "customer_id" VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "balance" FLOAT NOT NULL,
  "acc_limit" FLOAT NOT NULL,
  "acc_reversal" FLOAT NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);