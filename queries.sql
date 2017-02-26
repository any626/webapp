CREATE EXTENSION citext;

CREATE TABLE users (
	id bigserial PRIMARY KEY,
	email citext UNIQUE NOT NULL,
	password varchar NOT NULL,
	first_name varchar(25),
	last_name varchar(25),
	created_at timestamptz DEFAULT NOW(),
	updated_at timestamptz DEFAULT NOW()
);