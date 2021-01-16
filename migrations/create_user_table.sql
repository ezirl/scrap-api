CREATE TABLE user
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    email    varchar(255) not null,
    password varchar(255) not null,
    token    varchar(255) not null,
    tariff   varchar(150) DEFAULT 1,
    requests int     DEFAULT 0
);

CREATE INDEX i_token
    ON user (token);

INSERT INTO user (email, password, token, tariff, requests) values ('shvedi@com.ru', '123123', '7d58b345d95b5b021dc16fd5b48a8f8b', 'enterprise', 0);

INSERT INTO user (email, password, token, tariff, requests) values ('07sima07@gmail.com', '123123', 'qwe123', 'enterprise', 0);
