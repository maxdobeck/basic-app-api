CREATE TABLE members(
	name text CONSTRAINT name_present NOT NULL,
	email text CONSTRAINT email_present NOT NULL UNIQUE,
	password text CONSTRAINT password_present NOT NULL,
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4()
);
