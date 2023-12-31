-- migrate:up
CREATE TABLE answers
(
    id              serial          PRIMARY KEY,
    user_id         int,
    card_id         int,
    points          int,
    CONSTRAINT answers_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT answers_card_id FOREIGN KEY (card_id) REFERENCES cards(id) ON DELETE CASCADE
);


-- migrate:down
DROP TABLE answers;
