
-- +migrate Up
CREATE TABLE customer (
  id bigserial NOT NULL,
  first_name varchar(32) NOT NULL,
  last_name varchar(32) NOT NULL,
  age int NOT NULL,
  nickname varchar(64),
  memo text,
  is_active boolean NOT NULL DEFAULT true,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE customer;
