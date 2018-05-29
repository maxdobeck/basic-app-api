CREATE TABLE enrollments(
	schedule_id uuid REFERENCES schedules,
	member_id uuid REFERENCES members,
	admin BOOLEAN CONSTRAINT admin_powers NOT NULL,
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4()
);
