create table Login(
	login_id serial not null primary key,
	login varchar(32),
	password varchar(32)
);