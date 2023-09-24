-- migrate:up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users
(
    id              serial          PRIMARY KEY,
    telegram_id     int         ,
    name            citext,
    language        varchar(50) ,
    answered_cards  jsonb
);

CREATE TYPE card_category AS ENUM ('movie');
CREATE TABLE cards
(
    id              serial          PRIMARY KEY,
    category        card_category   NOT NULL,
    name_ru         citext,
    name_en         citext,
    desc_ru         varchar(250),
    desc_en         varchar(250),
    quote_ru        varchar(250),
    quote_en        varchar(250),
    facts_ru        jsonb,
    facts_en        jsonb,
    answers         jsonb,
    drawing_url     varchar(250),
    pixelated_url   varchar(250),
    screenshot_url  varchar(250),
    bg_url          varchar(250)
);



-- migrate:down
DROP EXTENSION IF EXISTS citext;

DROP TABLE users;

DROP TYPE card_category;

DROP TABLE cards;