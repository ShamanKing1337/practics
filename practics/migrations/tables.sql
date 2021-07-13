create table users (
    id bigserial primary key,
    login text Not null ,
    password text,
    last_name text not null,
    first_name text
);

create table items (
    id bigserial primary key,
    title text not null,
    description text,
    cost bigint not null
);

create table purchases (
    id bigserial primary key,
    user_id bigint not null,
    item_id bigint not null,
    purchase_time timestamptz not null
);

create table returned_purchases (
    id bigint not null,
    user_id bigint not null,
    item_id bigint not null,
    purchase_time timestamptz not null,
    return_time timestamptz not null
                                );