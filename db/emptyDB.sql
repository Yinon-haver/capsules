DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS capsules_users;
DROP TABLE IF EXISTS capsules;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
	phone CHAR(10) PRIMARY KEY NOT NULL,
	token CHAR(10) NOT NULL
);

CREATE TABLE capsules (
	id SERIAL PRIMARY KEY NOT NULL,
	from_phone CHAR(10) NOT NULL,
	posted_on TIMESTAMP NOT NULL,
	opened_on TIMESTAMP NOT NULL,
	FOREIGN KEY (from_phone) REFERENCES users(phone)
);

CREATE TABLE capsules_users (
	capsule_id INT,
	user_phone CHAR(10) NOT NULL,
	is_watched BOOL,
	FOREIGN KEY (capsule_id) REFERENCES capsules(id),
	FOREIGN KEY (user_phone) REFERENCES users(phone),
	PRIMARY KEY(capsule_id, user_phone)
);

CREATE TABLE messages (
	capsule_id INT,
	message_date TIMESTAMP NOT NULL,
	from_user CHAR(10) NOT NULL,
	content TEXT NOT NULL,
	FOREIGN KEY (capsule_id) REFERENCES capsules(id),
	FOREIGN KEY (from_user) REFERENCES users(phone),
	PRIMARY KEY(capsule_id, message_date)
);











