create table if not exists users(
    user_id serial primary key,
    user_name VARCHAR(16) not null,
    user_email VARCHAR(255) not null,
    user_passwd VARCHAR(255) not null,
    user_gender int not null
);

create table if not exists requests(
    request_number serial primary key,
    title VARCHAR(64) not null,
    description VARCHAR(255) not null,
    user_id int not null,
    photo VARCHAR(255)
);