
-- +migrate Up
CREATE TABLE scripts (
  id SERIAL PRIMARY KEY,
  name text NOT NULL,
  category varchar(20) NOT NULL,
  content text NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);
-- +migrate Down
DROP TABLE scripts;
