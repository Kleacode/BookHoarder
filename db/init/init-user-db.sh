#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username ${POSTGRES_USER} --dbname ${POSTGRES_DB} <<-EOSQL
	CREATE TABLE users(
	  id        SERIAL PRIMARY KEY,

	  name      VARCHAR(20)
	);
	INSERT INTO users (name) VALUES ('inituser1');
	INSERT INTO users (name) VALUES ('inituser2');
	INSERT INTO users (name) VALUES ('inituser3');

	CREATE TABLE books(
	  id        SERIAL PRIMARY KEY,

	  title     VARCHAR(100),
	  user_id	INT REFERENCES users(id) NOT NULL,
	  tags_id	INT[]
	);
	INSERT INTO books (title, user_id) VALUES ('initbook', 1);

	CREATE TABLE status(
	  id        SERIAL PRIMARY KEY,

	  name      VARCHAR(20)
	);
	INSERT INTO status (name) VALUES ('todo');
	INSERT INTO status (name) VALUES ('wip');
	INSERT INTO status (name) VALUES ('done');

	CREATE TABLE tags(
	  id        SERIAL PRIMARY KEY,

	  name      VARCHAR(50),
	  user_id	INT REFERENCES users(id) NOT NULL
	);
	INSERT INTO tags (name, user_id) VALUES ('tag1', 1);
	INSERT INTO tags (name, user_id) VALUES ('tag2', 1);
	INSERT INTO tags (name, user_id) VALUES ('tag3', 1);

	CREATE TABLE user_book_status(
	  id        SERIAL PRIMARY KEY,

	  user_id	INT REFERENCES users(id) NOT NULL,
	  book_id	INT REFERENCES books(id) NOT NULL,
	  status_id	INT REFERENCES status(id) NOT NULL
	);

EOSQL
