CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE schedules(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  title text CONSTRAINT title_present NOT NULL,
  sunday BOOLEAN,
  monday BOOLEAN,
  tuesday BOOLEAN,
  wednesday BOOLEAN,
  thursday BOOLEAN,
  friday BOOLEAN,
  saturday BOOLEAN
);

CREATE TABLE members(
	name text CONSTRAINT name_present NOT NULL,
	email text CONSTRAINT email_present NOT NULL UNIQUE,
	password text CONSTRAINT password_present NOT NULL,
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4()
);


CREATE TABLE enrollments(
	schedule_id uuid REFERENCES schedules,
	member_id uuid REFERENCES members,
	admin BOOLEAN CONSTRAINT admin_powers NOT NULL,
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4()
);

