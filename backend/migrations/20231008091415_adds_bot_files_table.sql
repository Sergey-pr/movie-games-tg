-- migrate:up
CREATE TABLE bot_files
(
    id          serial          PRIMARY KEY,
    filename    varchar(250),
    file_id     varchar(250)
);

-- migrate:down
DROP TABLE bot_files;
