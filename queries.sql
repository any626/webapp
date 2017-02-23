CREATE EXTENSION citext;

CREATE TABLE users (
	id bigserial PRIMARY KEY,
	email citext UNIQUE,
	password varchar NOT NULL,
	first_name varchar(25) NOT NULL,
	last_name varchar(25) NOT NULL,
	created_at timestamptz DEFAULT NOW(),
	updated_at timestamptz DEFAULT NOW()
);