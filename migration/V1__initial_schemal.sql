CREATE TABLE types (
	id integer NOT NULL ,
	name varchar(255) NOT NULL,
	img json DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
    PRIMARY KEY ( id )
);

CREATE TABLE stores  (
	id integer NOT NULL ,
	owner_id integer NOT NULL,
	type_id integer NOT NULL,
	name varchar(255) NOT NULL,
	description text DEFAULT NULL,
	cover json DEFAULT NULL, 
	logo json DEFAULT NULL,
	lng double precision DEFAULT NULL,
	lat double precision DEFAULT NULL,
	address varchar(255) DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
    PRIMARY KEY ( id )
);

CREATE TABLE  closed_date (
	id integer NOT NULL ,
	store_id integer NOT NULL,
	start_date_time timestamp NOT NULL,
	end_date_time timestamp NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY ( id , store_id )
);

CREATE TABLE  business_time  (
	id integer NOT NULL ,
	store_id integer NOT NULL,
	start_time timestamp NOT NULL,
	end_time timestamp NOT NULL,
	day_of_week varchar(10) NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( id , store_id )
);
INSERT INTO business_time (id, store_id, start_time, end_time, day_of_week, created_at, updated_at)
VALUES ('1', '1','2021-09-24 07:00:00','2021-09-24 17:00:00', 'MONDAY','2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO business_time (id, store_id, start_time, end_time, day_of_week, created_at, updated_at)
VALUES ('2', '1','2021-09-24 07:00:00','2021-09-24 17:00:00', 'TUESDAY','2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO business_time (id, store_id, start_time, end_time, day_of_week, created_at, updated_at)
VALUES ('3', '1','2021-09-24 07:00:00','2021-09-24 17:00:00', 'WEDNESDAY','2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO business_time (id, store_id, start_time, end_time, day_of_week, created_at, updated_at)
VALUES ('4', '1','2021-09-24 07:00:00','2021-09-24 17:00:00', 'THURSDAY','2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO business_time (id, store_id, start_time, end_time, day_of_week, created_at, updated_at)
VALUES ('5', '1','2021-09-24 07:00:00','2021-09-24 17:00:00', 'FRIDAY','2021-09-24 09:16:08', '2021-09-24 09:16:08');
 
CREATE TABLE  users  (
	id integer NOT NULL ,
	fb_id varchar(50) DEFAULT NULL ,
	gg_id varchar(50) DEFAULT NULL ,
	email varchar(50) NOT NULL ,
	password varchar(50) NOT NULL ,
	salt varchar(50) DEFAULT NULL ,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	phone varchar(20) DEFAULT NULL , 
	avatar json DEFAULT NULL,
	lat double precision DEFAULT NULL ,
	lng double precision  DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY ( id )
);
INSERT INTO users (id, fb_id , gg_id , email, password, salt, first_name, last_name, phone, avatar,lat,lng ,status, created_at, updated_at)
VALUES ('1',null,null,'user@gmail.com','3d55a249c559fdd849bc80f07022d659','uEbHhwJMLSsxknTbvFGEqvygxtEUcqxGWUXAqeHeerxjGzPwEk',
'user','user','0184632159',null,null,null,1,'2021-09-24 09:16:08', '2021-09-24 09:16:08');

CREATE TYPE Gender AS ENUM ('MALE', 'FEMALE');

CREATE TABLE  clients  (
	id integer NOT NULL ,
	first_name varchar(50) NOT NULL ,
	last_name varchar(50) NOT NULL ,
	gender Gender ,
	phone varchar(20) DEFAULT NULL , 
	email varchar(50) NOT NULL ,
	password varchar(50) NOT NULL ,
	salt varchar(50) DEFAULT NULL ,
	avatar json DEFAULT NULL,
	birthday Date DEFAULT NULL,
	address varchar(255) NOT NULL ,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY ( id )
);

CREATE TYPE Role AS ENUM (('owner','staff ','admin');

CREATE TABLE  staffs  (
	id integer NOT NULL ,
	store_id integer NOT NULL,
	first_name varchar(50) NOT NULL ,
	last_name varchar(50) NOT NULL ,
	phone varchar(20) DEFAULT NULL , 
	email varchar(50) NOT NULL ,
	password varchar(50) NOT NULL ,
	salt varchar(50) DEFAULT NULL ,
	role Role NOT NULL,
	avatar json DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
  PRIMARY KEY ( id )
);
INSERT INTO staffs (id, store_id, first_name, last_name, phone, email, password, salt, role , avatar, status, created_at, updated_at)
VALUES ('1', '1','admin','admin','0123456789','admin@gmail.com','3d55a249c559fdd849bc80f07022d659','uEbHhwJMLSsxknTbvFGEqvygxtEUcqxGWUXAqeHeerxjGzPwEk'
,'owner',null,1,'2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO staffs (id, store_id, first_name, last_name, phone, email, password, salt, role , avatar, status, created_at, updated_at)
VALUES ('2', '1','staff','staff','0987654321','staff@gmail.com','04f6cbedf46e746614916262ff2eaa10','ZnGmaNTaCVTIgTQAYTBhoxlwhTwQiIDBtuDwrKHZUYYujNrVlm'
,'staff',null,1,'2021-09-24 09:16:08', '2021-09-24 09:16:08');

CREATE TABLE  services  (
	id integer NOT NULL ,
	store_id integer NOT NULL,
	category_id integer NOT NULL ,
	title varchar(100) DEFAULT NULL ,
	description text DEFAULT NULL ,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( id )
);
INSERT INTO services (id, store_id, category_id, title, description, status, created_at, updated_at)
VALUES ('1', '1','1','Massage',null,1,'2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO services (id, store_id, category_id, title, description, status, created_at, updated_at)
VALUES ('2', '1','1','Lose Weight',null,1,'2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO services (id, store_id, category_id, title, description, status, created_at, updated_at)
VALUES ('3', '1','1','Facial Skin Care',null,1,'2021-09-24 09:16:08', '2021-09-24 09:16:08');


CREATE TABLE  category  (
	id integer NOT NULL ,
	name varchar(50),
	description text,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( id )
);

CREATE TABLE  service_details   (
	id integer NOT NULL ,
	service_id integer NOT NULL,
	title varchar(100) DEFAULT NULL ,
	price float  NOT NULL,
	duration int NOT NULL,
	description text DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( id , service_id )
);
 
CREATE TABLE  service_staffs  (
	service_id integer NOT NULL,
	staff_id integer NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( staff_id , service_id )
);

DROP TABLE IF EXISTS  staff_working ; 
CREATE TABLE  staff_working   (
	id integer NOT NULL,
	staff_id integer NOT NULL,
	start_time timestamp NOT NULL,
	end_time timestamp NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( staff_id , id )
);
INSERT INTO staff_working (id, staff_id, start_time, end_time,created_at,updated_at)
VALUES ('1', '1', '2021-09-24 07:00:00', '2021-09-24 15:50:00','2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO staff_working (id, staff_id, start_time, end_time ,created_at,updated_at)
VALUES ('2', '2', '2021-09-24 09:00:00', '2021-09-24 09:50:00','2021-09-24 09:16:08', '2021-09-24 09:16:08');
 
CREATE TABLE  staff_working_repeat  (
	id integer NOT NULL,
	staff_id integer NOT NULL,
	start_time timestamp NOT NULL,
	end_time timestamp NOT NULL,
	specific_date Date NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( staff_id , id )
);
 
CREATE TABLE  appointments  (
	id integer NOT NULL ,
	user_id integer DEFAULT NULL,
	store_id integer DEFAULT NULL,
	staff_id integer NOT NULL,
	start_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	end_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	total_price float NOT NULL,
	status int NOT NULL DEFAULT '1',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( id )
);
INSERT INTO appointments (id, user_id, store_id, staff_id, start_time, end_time, total_price,status,created_at,updated_at)
VALUES ('1', '1', '1', '1', '2021-09-24 09:00:00', '2021-09-24 09:50:00', '200', '1', '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointments (id, user_id, store_id, staff_id, start_time, end_time, total_price,status,created_at,updated_at)
VALUES ('2', '1', '1', '1', '2021-09-24 07:00:00', '2021-09-24 08:50:00', '200', '1', '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointments (id, user_id, store_id, staff_id, start_time, end_time, total_price,status,created_at,updated_at)
VALUES ('3', '2', '1', '2', '2021-09-24 13:00:00', '2021-09-24 15:50:00', '200', '1', '2021-09-24 09:16:08', '2021-09-24 09:16:08');

CREATE TABLE  appointment_service (
	service_id integer NOT NULL,
	appointment_id integer NOT NULL,
	status int NOT NULL DEFAULT '1',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( appointment_id , service_id )
);
INSERT INTO appointment_service (service_id,appointment_id,status,created_at,updated_at)
VALUES (1, 1 , 1, '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointment_service (service_id,appointment_id,status,created_at,updated_at)
VALUES (2, 1 , 1, '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointment_service (service_id,appointment_id,status,created_at,updated_at)
VALUES (1, 2 , 1, '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointment_service (service_id,appointment_id,status,created_at,updated_at)
VALUES (1, 3 , 1, '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointment_service (service_id,appointment_id,status,created_at,updated_at)
VALUES (2, 3 , 1, '2021-09-24 09:16:08', '2021-09-24 09:16:08');
INSERT INTO appointment_service (service_id,appointment_id,status,created_at,updated_at)
VALUES (3, 3 , 1, '2021-09-24 09:16:08', '2021-09-24 09:16:08');

CREATE TABLE  store_rating (
	user_id int NOT NULL,
	store_id int NOT NULL,
	point float NOT NULL,
	cmt text DEFAULT NULL,
	status int DEFAULT 1,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ( user_id , store_id ),
  KEY  store_id  ( store_id )
);

