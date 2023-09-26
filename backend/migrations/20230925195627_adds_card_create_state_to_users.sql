-- migrate:up
ALTER TABLE users add column card_create_state int default 0;

-- migrate:down
ALTER TABLE users drop column card_create_state;
