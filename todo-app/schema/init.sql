CREATE TABLE IF NOT EXISTS t_users (
    id serial Primary key,
    username varchar(50) unique not null,
    password varchar(50) not null, 
    email varchar(255) not null, 
    name varchar(50) not null, 
    surname varchar(50) not null
);

