create role postgres login password 'postgres';
create database postgres;
grant all privileges on database postgres to postgres;

CREATE TABLE TEST_USER (
    user_id BIGINT PRIMARY KEY,
    user_password VARCHAR(20) NOT NULL
);