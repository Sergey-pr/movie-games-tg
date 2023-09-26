-- migrate:up
ALTER TABLE cards RENAME COLUMN answers TO answers_en;
ALTER TABLE cards ADD COLUMN answers_ru JSONB;

-- migrate:down
ALTER TABLE cards RENAME COLUMN answers_en TO answers;
ALTER TABLE cards DROP COLUMN answers_ru;