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