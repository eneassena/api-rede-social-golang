CREATE DATABASE devbook;
USE devbook;

drop table if exists usuarios ;

create table usuarios(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) not null,
    nick VARCHAR(50) not null UNIQUE,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;
