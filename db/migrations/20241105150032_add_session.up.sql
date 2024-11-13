CREATE TABLE user_sessions (
  id UUID NOT NULL UNIQUE,
  user_id bigint NOT NULL,
  refresh_token varchar(150) NOT NULL,
  user_agent varchar(150) NOT NULL,
  client_ip varchar(150) NOT NULL,
  is_blocked boolean NOT NULL  DEFAULT false,
  expired_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY(id),
  FOREIGN kEY(user_id) REFERENCES users(id)
  ) engine = InnoDB;