package config

const TABLE_USERS = `CREATE TABLE IF NOT EXISTS users(
	id int auto_increment,
	phone varchar(12) NOT NULL UNIQUE,
	password text NOT NULL,
	name varchar(50) NOT NULL,
	date_of_birth DATE NOT NULL,
	sex ENUM("Pria","Wanita")),
	created_at datetime default current_timestamp(),
	updated_at datetime default current_timestamp(),
	deleted_at datetime
)`

const TABLE_BALANCES = `CREATE TABLE IF NOT EXISTS balances(
	user_id int not null auto_increment,
	balance decimal not null,
	created_at datetime default current_timestamp(),
	updated_at datetime default current_timestamp(),
	constraint fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
)`

const TABLE_TOP_UP_HISTORIES = `CREATE TABLE IF NOT EXISTS top_up_histories(
	id int auto_increment,
	user_id int not null,
	date datetime default current_timestamp(),
	amount decimal not null,
	info varchar(250),
	constraint fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
)`

const TABLE_TRANSFER_HISTORIES = `CREATE TABLE IF NOT EXISTS transfer_histories(
	id int auto_increment,
	user_id_sender int not null,
	user_id_receiver int not null,
	date datetime default current_timestamp(),
	amount decimal not null,
	info varchar(250),
	constraint fk_user_id_sender FOREIGN KEY (user_id_sender) REFERENCES users(id),
	constraint fk_user_id_receiver FOREIGN KEY (user_id_receiver) REFERENCES users(id)
)`
