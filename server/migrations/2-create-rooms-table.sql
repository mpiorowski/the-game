-- +migrate Up
CREATE TABLE
  rooms (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    created timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted timestamptz,
    name text UNIQUE NOT NULL,
    password text NOT NULL
  );

CREATE TRIGGER set_timestamp BEFORE
UPDATE
  ON rooms FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp ();

