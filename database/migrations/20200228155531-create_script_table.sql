
-- +migrate Up
CREATE TABLE scripts (
  id int PRIMARY KEY,
  name text NOT NULL,
  category varchar(20) NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);
-- +migrate Down
DROP TABLE scripts;
