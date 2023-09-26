-- migrate:up
ALTER TABLE users add column username varchar(250);
ALTER TABLE users add column is_admin bool default false;

-- migrate:down
ALTER TABLE users drop column username;
ALTER TABLE users drop column is_admin;
