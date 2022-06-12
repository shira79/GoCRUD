-- +migrate Up
CREATE TABLE IF NOT EXISTS  articles (
    id int AUTO_INCREMENT,
    title varchar(100),
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE IF EXISTS articles;
