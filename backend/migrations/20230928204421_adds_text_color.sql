-- migrate:up
ALTER TABLE cards ADD COLUMN text_color varchar(10);

-- migrate:down
ALTER TABLE cards DROP COLUMN IF EXISTS text_color;
