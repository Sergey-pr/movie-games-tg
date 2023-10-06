-- migrate:up
ALTER TABLE cards ALTER COLUMN desc_en TYPE varchar(1000);
ALTER TABLE cards ALTER COLUMN desc_ru TYPE varchar(1000);


-- migrate:down
ALTER TABLE cards ALTER COLUMN desc_en TYPE varchar(500);
ALTER TABLE cards ALTER COLUMN desc_ru TYPE varchar(500);

