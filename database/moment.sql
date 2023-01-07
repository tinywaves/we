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

select m.id      as                                      id,
       m.content as                                      content,
       m.createAt                                        createTime,
       m.updateAt                                        updateTime,
       JSON_OBJECT('id', u.id, 'authorName', u.username) author
from moment m
         left join users u on u.id = m.user_id
where m.id = 11;
