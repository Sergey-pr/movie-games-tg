-- migrate:up
alter table users alter column telegram_id type bigint;

-- migrate:down
alter table users drop column telegram_id;
