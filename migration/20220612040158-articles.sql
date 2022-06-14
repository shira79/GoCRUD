-- +migrate Up
CREATE TABLE IF NOT EXISTS articles (
    id int AUTO_INCREMENT,
    title varchar(100) NOT NULL,
    body mediumtext NOT NULL,
    created datetime,
    updated datetime,
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE IF EXISTS articles;
