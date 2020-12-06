
-- +migrate Up
CREATE TABLE customer (
  id bigserial NOT NULL,
  first_name varchar(32) NOT NULL,
  last_name varchar(32) NOT NULL,
  age int NOT NULL,
  nickname varchar(64),
  memo text,
  PRIMARY KEY (id)
);

CREATE TABLE todo (
  id bigserial NOT NULL,
  task varchar(64) NOT NULL,
  has_deadline boolean NOT NULL DEFAULT false,
  user_id bigint NOT NULL REFERENCES customer(id),
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE customer;
DROP TABLE todo;
