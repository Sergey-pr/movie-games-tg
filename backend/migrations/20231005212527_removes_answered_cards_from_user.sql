-- migrate:up
ALTER TABLE users DROP COLUMN answered_cards;

-- migrate:down

ALTER TABLE users ADD COLUMN answered_cards jsonb;