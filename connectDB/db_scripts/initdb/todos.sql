CREATE TABLE IF NOT EXISTS  todos(
  id serial not null PRIMARY KEY
, title varchar(50)
, description varchar(50)
, status boolean not null default false
, user_id  integer not null references users (id)
, created_at timestamp with time zone NOT NULL
, updated_at timestamp with time zone NOT NULL
, deleted_at timestamp with time zone NULL
);