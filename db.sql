drop database if exists `test_db`;

create database test_db;
use test_db;

create table test (
    `id` int NOT NULL auto_increment,
    primary key(id),
    `name` varchar(255),
    `time` DATETIME
);

