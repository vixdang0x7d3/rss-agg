-- +goose UP

create table rssagg.users (
    id uuid primary key, 
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null
);


-- +goose DOWN
drop table rssagg.users; 
