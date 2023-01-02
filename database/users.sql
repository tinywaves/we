create table if not exists `users`
(
    id       int primary key auto_increment,
    username varchar(20) not null unique,
    password varchar(50) not null,
    createAt timestamp default current_timestamp,
    updateAt timestamp default current_timestamp on update current_timestamp
);

select *
from users;
