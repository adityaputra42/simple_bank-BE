CREATE TABLE transactions(
  id varchar(100) NOT NULL UNIQUE,
  from_account_id bigint NOT NULL,
  to_account_id bigint NOT NULL,
  amount bigint NOT NULL,
  currency varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(from_account_id) REFERENCES accounts(id),
  FOREIGN KEY(to_account_id)REFERENCES accounts(id)
  ) engine = InnoDB;