DROP TABLE IF EXISTS `types`; 
CREATE TABLE `types` (
	id int(10) NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	status int NOT NULL DEFAULT '1',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
);

DROP TABLE IF EXISTS `stores`;  
CREATE TABLE `stores`  (
	id int(10) NOT NULL AUTO_INCREMENT,
	owner_id int(10) NOT NULL,
	type_id int(10) NOT NULL,
	name varchar(255) NOT NULL,
	description text DEFAULT NULL,
	cover json DEFAULT NULL, 
	logo json DEFAULT NULL,
	lng double DEFAULT NULL,
	lat double DEFAULT NULL,
	address varchar(255) DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `type_id` (`type_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
);

DROP TABLE IF EXISTS `closed_date`;
CREATE TABLE `closed_date`(
	id int(10) NOT NULL AUTO_INCREMENT,
	store_id int(10) NOT NULL,
	start_date_time timestamp NOT NULL,
	end_date_time timestamp NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`store_id`)
);

DROP TABLE IF EXISTS `business_time`; 
CREATE TABLE `business_time` (
	id int(10) NOT NULL AUTO_INCREMENT,
	store_id int(10) NOT NULL,
	start_time time NOT NULL,
	end_time time NOT NULL,
	day_of_week enum('MONDAY', 'THURSDAY', 'WEDNESDAY', 'TUESDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY') NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`,`store_id`)
);

DROP TABLE IF EXISTS `users`; 
CREATE TABLE `users` (
	id int(10) NOT NULL AUTO_INCREMENT,
	fb_id varchar(50) DEFAULT NULL ,
	gg_id varchar(50) DEFAULT NULL ,
	email varchar(50) NOT NULL ,
	password varchar(50) NOT NULL ,
	salt varchar(50) DEFAULT NULL ,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	phone varchar(20) DEFAULT NULL , 
	avatar json DEFAULT NULL,
	lat Double DEFAULT NULL ,
	lng Double  DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);

DROP TABLE IF EXISTS `clients`; 
CREATE TABLE `clients` (
	id int(10) NOT NULL AUTO_INCREMENT,
	first_name varchar(50) NOT NULL ,
	last_name varchar(50) NOT NULL ,
	gender enum('MALE', 'FEMALE') ,
	phone varchar(20) DEFAULT NULL , 
	email varchar(50) NOT NULL ,
	password varchar(50) NOT NULL ,
	salt varchar(50) DEFAULT NULL ,
	avatar json DEFAULT NULL,
	birthday Date DEFAULT NULL,
	address varchar(255) NOT NULL ,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);

DROP TABLE IF EXISTS `staffs`; 
CREATE TABLE `staffs` (
	id int(10) NOT NULL AUTO_INCREMENT,
	store_id int(10) NOT NULL,
	first_name varchar(50) NOT NULL ,
	last_name varchar(50) NOT NULL ,
	phone varchar(20) DEFAULT NULL , 
	email varchar(50) NOT NULL ,
	password varchar(50) NOT NULL ,
	salt varchar(50) DEFAULT NULL ,
	role enum('owner','staff ','admin') NOT NULL,
	avatar json DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `store_id` (`store_id`) USING BTREE,
  UNIQUE KEY `email` (`email`)
);

DROP TABLE IF EXISTS `services`; 
CREATE TABLE `services` (
	id int(10) NOT NULL AUTO_INCREMENT,
	store_id int(10) NOT NULL,
	category_id int(10) DEFAULT NULL ,
	title varchar(100) DEFAULT NULL ,
	description text DEFAULT NULL ,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `store_id` (`store_id`) USING BTREE
);

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
	id int(10) NOT NULL AUTO_INCREMENT,
	name varchar(50),
	description text,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `service_details`; 
CREATE TABLE `service_details`  (
	id int(10) NOT NULL AUTO_INCREMENT,
	service_id int(10) NOT NULL,
	title varchar(100) DEFAULT NULL ,
	price float  NOT NULL,
	duration time NOT NULL,
	description text DEFAULT NULL,
	status int NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`,`service_id`)
);

DROP TABLE IF EXISTS `service_staffs`; 
CREATE TABLE `service_staffs` (
	service_id int(10) NOT NULL,
	staff_id int(10) NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`staff_id`,`service_id`)
);

DROP TABLE IF EXISTS `service_working`; 
CREATE TABLE `service_working`  (
	id int(10) NOT NULL AUTO_INCREMENT,
	staff_id int(10) NOT NULL,
	start_time time NOT NULL,
	end_time time NOT NULL,
	working_date Date NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`staff_id`,`id`)
);

DROP TABLE IF EXISTS `service_working_repeat`; 
CREATE TABLE `service_working_repeat` (
	id int(10) NOT NULL AUTO_INCREMENT,
	staff_id int(10) NOT NULL,
	start_time time NOT NULL,
	end_time time NOT NULL,
	working_date Date NOT NULL,
	specific_date Date NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`staff_id`,`id`)
);

DROP TABLE IF EXISTS `appointments`; 
CREATE TABLE `appointments` (
	id int(10) NOT NULL AUTO_INCREMENT,
	client_id int(10) DEFAULT NULL,
	service_id int(10) NOT NULL,
	user_id int(10) NOT NULL,
	staff_id int(10) NOT NULL,
	duration time NOT NULL,
	price float NOT NULL,
	status int NOT NULL DEFAULT '1',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`,`service_id`)
);

DROP TABLE IF EXISTS `appointment_service`; 
CREATE TABLE `appointment_service`(
	service_id int(10) NOT NULL,
	appointment_id int(10) NOT NULL,
	status int NOT NULL DEFAULT '1',
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`appointment_id`,`service_id`)
);

