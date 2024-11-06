-- Create the database and user
CREATE DATABASE mydb;
CREATE USER user WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE mydb TO user;

\c mydb

-- Create the sample table
CREATE TABLE IF NOT EXISTS sample_table (
    pkid SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
