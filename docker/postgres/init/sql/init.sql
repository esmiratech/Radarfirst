\set database_name test_db
\set schema_name hello
\set app_user db_user


CREATE DATABASE :database_name;

CREATE USER :app_user WITH PASSWORD 'app_user_password';

GRANT ALL PRIVILEGES ON DATABASE :database_name TO :app_user;

\c :database_name;

CREATE SCHEMA :schema_name;

CREATE TABLE :schema_name.world (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now()
);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA :schema_name TO :app_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA :schema_name TO :app_user;
GRANT USAGE ON SCHEMA :schema_name TO :app_user;
