-- migrate:up
ALTER TABLE users drop column card_create_state;

CREATE TABLE card_processors
(
    id      serial  PRIMARY KEY,
    user_id int,
    card_id int,
    state   int     default 0,
    CONSTRAINT card_processors_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT card_processors_cards_fk FOREIGN KEY (card_id) REFERENCES cards (id)
);

CREATE INDEX card_processors_users_fk ON card_processors (user_id);
CREATE INDEX card_processors_cards_fk ON card_processors (card_id);


-- migrate:down
ALTER TABLE users add column card_create_state int default 0;

DROP TABLE card_processors;
DROP INDEX card_processors_users_fk;
DROP INDEX card_processors_cards_fk;