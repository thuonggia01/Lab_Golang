CREATE TABLE IF NOT EXISTS  users(
  id serial not null PRIMARY KEY
, name varchar(50) NOT NULL
, todo_id integer NOT NULL references todos(id)
, created_at timestamp with time zone NOT NULL
, updated_at timestamp with time zone NOT NULL
, deleted_at timestamp with time zone NULL
);