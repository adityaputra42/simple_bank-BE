CREATE TABLE users (
  id bigint NOT NULL AUTO_INCREMENT,
  username varchar(100)  NOT NULL,
  password varchar(100) NOT NULL,
  full_name varchar(255) NOT NULL,
  email varchar(100) UNIQUE NOT NULL,
  updated_at timestamp NOT NULL DEFAULT current_timestamp  ON UPDATE current_timestamp,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)  
) engine = InnoDB;

