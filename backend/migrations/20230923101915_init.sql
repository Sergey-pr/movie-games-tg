-- migrate:up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users
(
    id              serial          PRIMARY KEY,
    tg_id           int             NOT NULL,
    name            varchar(250)    NOT NULL,
    language        varchar(50)     NOT NULL,
    answered_cards  jsonb
);

CREATE TYPE card_category AS ENUM ('movie');
CREATE TABLE cards
(
    id              serial          PRIMARY KEY,
    category        card_category   NOT NULL,
    name_ru         varchar(250)    NOT NULL,
    name_en         varchar(250)    NOT NULL,
    desc_ru         varchar(250)    NOT NULL,
    desc_en         varchar(250)    NOT NULL,
    quote_ru        varchar(250)    NOT NULL,
    quote_en        varchar(250)    NOT NULL,
    facts_ru        jsonb           NOT NULL,
    facts_en        jsonb           NOT NULL,
    answers         jsonb           NOT NULL,
    drawing_url     varchar(250)    NOT NULL,
    pixelated_url   varchar(250)    NOT NULL,
    screenshot_url  varchar(250)    NOT NULL,
    bg_url          varchar(250)    NOT NULL
);



-- migrate:down
DROP EXTENSION IF EXISTS citext;

DROP TABLE users;

DROP TYPE card_category;

DROP TABLE cards;