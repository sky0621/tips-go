CREATE SEQUENCE IF NOT EXISTS reviews_id_seq;

CREATE TABLE IF NOT EXISTS seller (
  company_id bigint NOT NULL,
  id         bigserial NOT NULL,
  name       character varying(255) NOT NULL,
  PRIMARY KEY (company_id, id)
);
INSERT INTO seller VALUES (1, 1, '売り主A');

CREATE TABLE IF NOT EXISTS contents (
  id   bigserial NOT NULL,
  name character varying(255) NOT NULL,
  PRIMARY KEY (id)
);
INSERT INTO contents VALUES (1, 'コンテンツA');

CREATE TABLE IF NOT EXISTS reviews (
  content_id bigint REFERENCES contents(id) NOT NULL,
  seller_id  bigint NOT NULL,
  reviews_id bigint DEFAULT nextval('reviews_id_seq'),
  result     integer NOT NULL,
  UNIQUE(reviews_id),
  PRIMARY KEY (content_id, seller_id)
);
INSERT INTO reviews(content_id, seller_id, result) VALUES (1, 1, 1);
