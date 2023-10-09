-- migrate:up
ALTER TABLE users ADD COLUMN last_name varchar(250);

-- migrate:down
ALTER TABLE users DROP COLUMN last_name;
