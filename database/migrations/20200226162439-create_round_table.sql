
-- +migrate Up
CREATE TABLE rounds (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    description text NOT NULL,
    users bigint NOT NULL,
    ramp_up bigint NOT NULL,
    duration bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
-- +migrate Down
DROP TABLE rounds;
