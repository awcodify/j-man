
-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name text,
  last_name text,
  email text not null unique,
  password_digest text not null
);

-- +migrate Down
DROP TABLE users;
