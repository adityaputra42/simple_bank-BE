CREATE TABLE accounts (
  id bigint NOT NULL AUTO_INCREMENT,
  user_id bigint NOT NULL,
  balance bigint NOT NULL,
  currency varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(user_id) REFERENCES users(id)
  ) engine = InnoDB;