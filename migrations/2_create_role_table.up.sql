create sequence tbl_role_seq;

create table tbl_role
(
    id           bigint default nextval ('tbl_role_seq')
        primary key,
    name     varchar(128) not null
);
