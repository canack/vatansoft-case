create table todos
(
    uuid       varchar(36)   not null,
    title      varchar(128)  not null,
    content    varchar(2048) null,
    created_at datetime      not null,
    updated_at datetime      not null,
    completed  bool          not null,
    priority   smallint default 0 not null,
    constraint todos_pk
        primary key (uuid)
)
    comment 'Creates todo table';

create unique index todos_uuid_uindex
    on todos (uuid);

create unique index todos_uuid_uindex_2
    on todos (uuid);

