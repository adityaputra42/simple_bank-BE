CREATE TABLE deposits(
  id bigint NOT NULL AUTO_INCREMENT,
  account_id bigint NOT NULL,
  amount bigint NOT NULL,
  currency varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY(id),
  FOREIGN KEY(account_id) REFERENCES accounts(id)

  ) engine = InnoDB;