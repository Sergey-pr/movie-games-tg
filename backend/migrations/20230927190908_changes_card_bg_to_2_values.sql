-- migrate:up
ALTER TABLE cards RENAME COLUMN bg_id TO bg_color_1;
ALTER TABLE cards ADD COLUMN bg_color_2 varchar(250);

-- migrate:down
ALTER TABLE cards RENAME COLUMN bg_color_1 TO bg_id;
ALTER TABLE cards DROP COLUMN bg_color_2;
