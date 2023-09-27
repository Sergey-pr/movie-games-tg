-- migrate:up
ALTER TABLE cards ADD COLUMN completed BOOLEAN DEFAULT false;

-- migrate:down
ALTER TABLE cards DROP COLUMN completed;
