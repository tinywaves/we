create table if not exists `moment`
(
    id       int primary key auto_increment,
    content  varchar(1000) not null,
    user_id  int           not null,
    createAt timestamp default current_timestamp,
    updateAt timestamp default current_timestamp on update current_timestamp,
    foreign key (user_id) references users (id)
);

select *
from moment;
