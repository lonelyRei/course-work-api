create table users
(
    id serial not null primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    about text not null default 'Я обычный пользователь сайта',
    image text not null default 'data:image/png;base64,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'
);

create table posts
(
    id serial not null primary key,
    user_id int not null,
    foreign key (user_id) references users(id) on delete cascade,
    title varchar(255) not null,
    content text not null,
    image text,
    rating int not null default 0
);