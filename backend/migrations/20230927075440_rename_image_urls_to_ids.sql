-- migrate:up
ALTER TABLE cards RENAME COLUMN drawing_url TO drawing_id;
ALTER TABLE cards RENAME COLUMN pixelated_url TO pixelated_id;
ALTER TABLE cards RENAME COLUMN screenshot_url TO screenshot_id;
ALTER TABLE cards RENAME COLUMN bg_url TO bg_id;
-- migrate:down
ALTER TABLE cards RENAME COLUMN drawing_id TO drawing_url;
ALTER TABLE cards RENAME COLUMN pixelated_id TO pixelated_url;
ALTER TABLE cards RENAME COLUMN screenshot_id TO screenshot_url;
ALTER TABLE cards RENAME COLUMN bg_id TO bg_url;