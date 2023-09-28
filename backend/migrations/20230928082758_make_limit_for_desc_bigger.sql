-- migrate:up
ALTER TABLE cards ALTER COLUMN desc_en TYPE varchar(500);
ALTER TABLE cards ALTER COLUMN desc_ru TYPE varchar(500);


-- migrate:down
ALTER TABLE cards ALTER COLUMN desc_en TYPE varchar(250);
ALTER TABLE cards ALTER COLUMN desc_ru TYPE varchar(250);

