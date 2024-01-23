create table register_user_step_one 
(
	"id" serial primary key not null,
	registration_id text unique not null,
	email text not null,
	constraint proper_email check (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')
);

create table register_user_step_two
(
	"id" serial primary key not null,
	registration_id text unique not null,
	verification_code text not null
);

create table register_user_step_three
(
	"id" serial primary key not null,
	registration_id text unique not null,
	firstname varchar(20) not null,
	lastname varchar(20) not null
);

create table register_user_step_four
(
	"id" serial primary key not null,
	registration_id text unique not null,
	username varchar(20) unique not null,
	constraint len_username check(length(username) > 3)
);

create table messanger_user
(
	"id" serial primary key not null,
	keycloak_id text unique not null,
	email text not null,
	firstname varchar(20) default null,
	lastname varchar(20) default null,
	username varchar(20) unique not null,
	is_blocker boolean not null,
	created_at timestamp not null,
	deleted_at timestamp default null,
	constraint proper_email check (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
	constraint len_username check(length(username) >= 3)
);

create table verification_email 
(
	"id" serial primary key not null,
	email text not null,
	verification_code text not null,
	created_at timestamp not null,
	constraint proper_email check (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')
);

create table chat 
(
	"id" serial primary key not null,
	id_first_user text not null,
	id_second_user text not null,
	created_at timestamp not null,
	deleted_at timestamp null,
	foreign key (id_first_user) references messanger_user (keycloak_id),
	foreign key (id_second_user) references messanger_user (keycloak_id)
);