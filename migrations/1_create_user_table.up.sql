create sequence tbl_user_seq;

create table tbl_user
(
    id           bigint default nextval ('tbl_user_seq')
        primary key,
    password     varchar(128) not null,
    last_login   timestamp(6)  null,
    is_superuser smallint   not null,
    email        varchar(254) not null,
    user_name    varchar(150) null,
    first_name   varchar(150) not null,
    start_date   timestamp(6)  not null,
    about        text     not null,
    image        varchar(100) null,
    is_active    smallint   not null,
    is_staff     smallint   not null,
    unique (email)
);
